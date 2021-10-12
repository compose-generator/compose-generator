## GitLab
GitLab is a full DevOps platform for managing, deploying and reviewing code in a single application.

### Setup
After setting up GitLab with Compose Generator, you can access it via port 80 / 443 and login with the username `root`. You can get the initial password for the user `root`, by executing `sudo docker exec -it ${{PROJECT_NAME_CONTAINER}}-frontend-gitlab grep 'Password:' /etc/gitlab/initial_root_password`.