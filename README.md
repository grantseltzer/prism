# Prism
Tool for analyzing the effects an application will have on your system when run

### Prism is still a work in progress

- It started as a hackathon project, so there's definitely parts that need revision/optimization

- Take a look at the issues if you'd like to contribute. (Thanks in advance!)


### The following steps are used to analyze an executable:

- A 'snapshot' is created of your linux filesystem. This includes path, last access timestamp, permission bits, and content (mocked).

- The snapshot and the executable is loaded into a container image.

- The container is run, and then the snapshot is used to repopulate your directory tree. 

- The executable is run.

- Another snapshot is taken.

- Compare the two snapshots and report back to user what was changed.

### What still needs to be done:

- JSON snapshot comparison.

- Improve bash script quality.

- Expand snapshot to include more than just filesystem objects.
