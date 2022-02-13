# Demo Application

The purpose of this demo app is to download a json file via an api, then back up to a git repo designated automatically. 

Later adding custom configuration of git repo and json file to download.

#### TODO Items

* Add capability to push to github branch specified by user
* Enable config file option to load user specific api keys, files, urls, etc.


#### Completed Items
* Logging to JSON file using [Zerolog](https://github.com/rs/zerolog "Zerolog")
* Added mapping for URLs and output files so a loop can be done on all links in a map with respective output files
