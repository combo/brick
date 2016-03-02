# Fields

* `name` (required): Package name.

    Shall begin and end with `[0-9a-z]` and optionally contain non-consecutive `[-_\.]`

    Valid: foo, foo-bar, foo_bar, foo.bar
    Invalid: foo_, .foo, foo--bar, foo__bar
