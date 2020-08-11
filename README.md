# Slide Window Practice

This is a slide window contains feature to find. Triple and fourth also the ordering number. Collected all possible results.

Here is the example:

```
example: []int{2, 2, 2, 2, 3, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 7, 8}

.
├── +[2 3 4]
│   ├── +[6 7 8]
│   ├── +[5 5 5]
│   │   ├── +[6 7 8]
│   │   ├── +[6 6 6]
│   │   │   └── +[6 7 8]
│   │   └── +[6 6 6 6]
│   │       └── +[6 7 8]
│   └── +[5 5 5 5]
│       ├── +[6 7 8]
│       ├── +[6 6 6]
│       │   └── +[6 7 8]
│       └── +[6 6 6 6]
│           └── +[6 7 8]
├── +[2 2 2 2]
│   ├── +[2 3 4]
│   │   ├── +[6 7 8]
│   │   ├── +[5 5 5]
│   │   │   ├── +[6 7 8]
│   │   │   ├── +[6 6 6]
│   │   │   │   └── +[6 7 8]
│   │   │   └── +[6 6 6 6]
│   │   │       └── +[6 7 8]
│   │   └── +[5 5 5 5]
│   │       ├── +[6 7 8]
│   │       ├── +[6 6 6]
│   │       │   └── +[6 7 8]
│   │       └── +[6 6 6 6]
│   │           └── +[6 7 8]
│   ├── +[5 5 5]
│   │   ├── +[6 7 8]
│   │   ├── +[6 6 6]
│   │   │   └── +[6 7 8]
│   │   └── +[6 6 6 6]
│   │       └── +[6 7 8]
│   └── +[5 5 5 5]
│       ├── +[6 7 8]
│       ├── +[6 6 6]
│       │   └── +[6 7 8]
│       └── +[6 6 6 6]
│           └── +[6 7 8]
└── +[2 2 2]
    ├── +[2 3 4]
    │   ├── +[6 7 8]
    │   ├── +[5 5 5]
    │   │   ├── +[6 7 8]
    │   │   ├── +[6 6 6]
    │   │   │   └── +[6 7 8]
    │   │   └── +[6 6 6 6]
    │   │       └── +[6 7 8]
    │   └── +[5 5 5 5]
    │       ├── +[6 7 8]
    │       ├── +[6 6 6]
    │       │   └── +[6 7 8]
    │       └── +[6 6 6 6]
    │           └── +[6 7 8]
    ├── +[5 5 5]
    │   ├── +[6 7 8]
    │   ├── +[6 6 6]
    │   │   └── +[6 7 8]
    │   └── +[6 6 6 6]
    │       └── +[6 7 8]
    └── +[5 5 5 5]
        ├── +[6 7 8]
        ├── +[6 6 6]
        │   └── +[6 7 8]
        └── +[6 6 6 6]
            └── +[6 7 8]
```
