Update 20180705 - awaiting fix from Apache Ignite team - https://issues.apache.org/jira/browse/IGNITE-8930 

# Description

Example app for reproducing the issue described here: https://github.com/alexbrainman/odbc/issues/121

Creates a single table, seeds with some hard-coded rows, then perform queries to replicate the issue 

# To replicate: 

* initiate local ignite cluster with `ignite-config.xml`
* compile and run go code
* Expected terminal output: 
```
...
...
...
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:05 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:06 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:06 Retrieved entry: {9661eef1-8cbd-4cc9-9eac-5786e137f803 1.3}
2018/06/29 20:46:06 SQLExecute: {HY000} Too many open cursors (either close other open cursors or increase the limit through ClientConnectorConfiguration.maxOpenCursorsPerConnection) [maximum=128, current=128]
```
