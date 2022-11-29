# gitctl

git control tool shows the PR data like:

input.json is the request param to this tool

Example
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
`{` 
    `"url":"https://api.github.com/repos/novuhq/novu",`
    `"prevDays":7,`
    `"sendermailID":"sender@gmail.com",`
    `"recievermailID":"receiver@gmail.com"`
`}`
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

1. url --> URL of git repository of which we want PR data
2. prevDays --> how much days we want data of PR
3. sendermailID --> email ID of sender user
4. recievermailID --> email ID of reciever user


# output:
 
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
Get PR data from github repo.....
To: user@gmail.com
From: gitctl@gmail.com
<<<<<<------- Here is the PR Data fro gitctl------->>>>>>
Subject: [DoNotReply] PR Report of last weeks github PRs for repo  novuhq/novu
`~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~`
 State of PR    ::      Count
`~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~`
 total    ->       49
 closed    ->       29
 merged    ->       24
 open    ->       20

 ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
