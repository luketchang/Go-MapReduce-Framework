# Go-MapReduce-Framework

### Summary
- **Overview**: Distributed MapReduce framework leveraging GCloud VMs for server and workers
- **Server**:
  - Reads arguments and supplied configuration file for initialization parameters (e.g. input/output directories, custom mapper/reducer scripts, number of mappers/reducers, etc)
  - Start listening on designated port
  - Spawns mappers, responding to mapper requests with new input files (from shared_files/input) and acknowledging mapper messages
  - Spawns reducers, responding to reducer requests with new intermediate files and acknowledging reducer messages
- **Mappers**:
  - Repeatedly requests input from server
  - Runs user-supplied custom mapper script on received input file
  - Buckets key-value pairs of mapped file into separate files based on key's hash-value (in shared_files/intermediate)
  - Notifies server of progress and final job status for each input
- **Reducers**:
  - Repeatedly requests intermediate files from server
  - Runs sort and 'group-by-key' script on received intermediat files
  - Runs user-supplied custom reducer script on sorted/grouped file (reduced files output to shared_files/output)
  - Notifies server of progress and final job status for each input
  
### Design Questions/Decisions
- Combined mapping and hash-bucketing steps under mapper but could have had 3 separate entities (mapper, bucketer, and reducer) to better adhere to Single Responsibility Principle
- Additionally, if bucketers were separate entities that ran concurrently with mappers, higher performance could be achieved with bucketers performing job as soon as new mapped file is available
- Used Google Filestore for shared storage across VMs but basic Filestore I/O ended up being extremely slow (would likely switch to different alternative or have workers send locally processed files back to server instead of having shared storage)

### Server Logs
<img src="https://raw.githubusercontent.com/ltchang2019/Go-MapReduce-Framework/master/log_images/mappers.png" width="500" height="400" />
<img src="https://raw.githubusercontent.com/ltchang2019/Go-MapReduce-Framework/master/log_images/reducers.png" width="500" height="400" />

### Todo
- *Add automated tests for verifying consistency of outputs of remote commands
- Add thorough error checking in argument and config parsers
- Do better job of surfacing error codes to server in functions called by remote executables
- Add builtin remote failures to test job rescheduling functionality
- Improve readability of util Contains function
- Fix network access and permissioning among Google Cloud VMs
- Switch from Filestore to different alternative for faster shared I/O
- Add user CLI functionality for uploading custom input and scripts to VMs
