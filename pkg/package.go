package pkg

import (
	"archive/tar"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const TREE_DIR = "tree"

// Add to tar tw all the files in root, stripping baseDir from the name in the archive
func tarRecursively(tw *tar.Writer, root, baseDir string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Build name
		name, err := filepath.Rel(baseDir, path)
		if err != nil {
			return err
		}

		// Write header
		hdr, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		hdr.Name = name

		if err = tw.WriteHeader(hdr); err != nil {
			return err
		}

		// Write contents
		if hdr.Size > 0 {
			input, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			if _, err = tw.Write(input); err != nil {
				return err
			}
		}

		return nil
	})
}

// Given a package directory containing metadata and tree, write the brick package to w
// Returns the new package metadata, and error, if any
func Package(w io.Writer, baseDir string) (*Metadata, error) {
	// Load the metadata file
	m, err := OpenMetadata(baseDir)
	if err != nil {
		return nil, err
	}

	// Validate its contents
	if err = m.Validate(); err != nil {
		return nil, err
	}

	// TODO: hashing

	// Create the package
	tw := tar.NewWriter(w)

	// Add the metadata file
	var mBuf bytes.Buffer

	if err = m.Write(&mBuf); err != nil {
		return nil, err
	}

	mHdr := &tar.Header{
		Name: METADATA_FILE,
		Mode: 0644,
		Size: int64(mBuf.Len()),
	}
	if err = tw.WriteHeader(mHdr); err != nil {
		return nil, err
	}
	if _, err := tw.Write(mBuf.Bytes()); err != nil {
		return nil, err
	}

	// Add the tree
	treeDir := filepath.Join(baseDir, TREE_DIR)
	if err = tarRecursively(tw, treeDir, baseDir); err != nil {
		return nil, err
	}

	return m, nil
}
