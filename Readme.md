## HOW TO RUN
   - Use `make build` to build the binary
   - Now Use ./treee and add `folder name` where you want use it 
      > eg 
           > - `` ./treee [folder name] `` for a specific folder
           > - `` ./tree `` for doing it in the cwd or root
   - Use `-f` for relative path
   - Use `-d` for tree only in directories
   - Use `-l` [value]` for printing nested levels upto the set [value] level
   - Use `-p` for printing file permissions for all file

## How to test
   - Use `make test` command for running the test
      - Current coverage is  `89.3%` for tree function(core)
      - Average `89.7%%`
