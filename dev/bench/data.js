window.BENCHMARK_DATA = {
  "lastUpdate": 1650465495805,
  "repoUrl": "https://github.com/posteris/commons",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "gsdenys@gmail.com",
            "name": "Denys G. Santos",
            "username": "gsdenys"
          },
          "committer": {
            "email": "gsdenys@gmail.com",
            "name": "Denys G. Santos",
            "username": "gsdenys"
          },
          "distinct": true,
          "id": "ed21f63fdb332d948aa8e49843c6aa9dd4773a12",
          "message": "add bench support",
          "timestamp": "2022-04-20T01:09:35-03:00",
          "tree_id": "911dff7c1e88cac53dc65671ef2d1875d41ce996",
          "url": "https://github.com/posteris/commons/commit/ed21f63fdb332d948aa8e49843c6aa9dd4773a12"
        },
        "date": 1650428260366,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkValidateModel/Errorless",
            "value": 0.0000441,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkValidateModel/Error",
            "value": 0.000202,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "gsdenys@gmail.com",
            "name": "Denys G. Santos",
            "username": "gsdenys"
          },
          "committer": {
            "email": "gsdenys@gmail.com",
            "name": "Denys G. Santos",
            "username": "gsdenys"
          },
          "distinct": true,
          "id": "5ea6031ffbc2c0abbe572e7843deaee56d4b0f0b",
          "message": "add benchmark",
          "timestamp": "2022-04-20T10:56:42-03:00",
          "tree_id": "84e39c04de0e014ee865af37e935098579de275f",
          "url": "https://github.com/posteris/commons/commit/5ea6031ffbc2c0abbe572e7843deaee56d4b0f0b"
        },
        "date": 1650465495340,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCreateDefaultError/with-message",
            "value": 0.0000014,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkCreateDefaultError/without-message",
            "value": 0.0000037,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkCreateValidationError/with-data",
            "value": 0.0000012,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkCreateValidationError/without-data",
            "value": 0.0000019,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkValidateModel/Errorless",
            "value": 0.0000076,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkValidateModel/Error",
            "value": 0.0001117,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          }
        ]
      }
    ]
  }
}