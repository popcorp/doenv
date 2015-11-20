# doenv
`doenv` is an easy way to manage your DigitalOcean instances in a Vagrant or virtualenv way.

## Overview
Docker and Vagrant are great, but heavy and quite limited for your development needs, unless you have infinite
computation power, optic fiber, open ports and so on…

On the other hand, deploying your development stack on the cloud might become expensive on the long run : if you work
on multiple projects and one of them is frozen, you still pay for it.

Well, forget it. With DigitalOcean you pay per hour. Oh, and when the project is frozen, you don't pay. Because your
instance is stopped and destroyed, but stored in a snapshot.

On the long run, starting new instances and stopping them is boring. Why can't we simply open our shell, type
"work on my-app", and start working on it ? And when we're done, simply typing "freeze my-app" to archive it.

That would be awesome. And that's what doenv does :stars: ! (What a surprise!)

## Installation
Make sure you have a working Go environment (go 1.5+ is recommanded). [See the install instructions](http://golang.org/doc/install.html).

To install `doenv`, simply run:
```
$ go get github.com/popcorp/doenv
```

When you launch doenv for the first time, it will ask for a Personnal API Token.
You might generate one [in the DigitalOcean Control Panel](https://cloud.digitalocean.com/settings/applications) (doenv requires the Write permission!).
Then, assuming the token is _62530fd95b6a48affc7bdeeccb01db0e2c09e18cac157f96f1a5e150e037d0aa_, run `doenv set-token 62530fd95b6a48affc7bdeeccb01db0e2c09e18cac157f96f1a5e150e037d0aa`. You're good to go !

If it tells you _doenv: command not found_, Make sure your `PATH` includes to the `$GOPATH/bin` directory so your commands can be easily used:
```
export PATH=$PATH:$GOPATH/bin
```

If you want to setup your SSH keys to be used, now looks like to be the best moment. doenv might use any [key that is setup on your DigitalOcean account](https://cloud.digitalocean.com/settings/security).
In my case, I have one key setup whose fingerprint is _1d:66:69:b5:d6:8e:0a:a6:4c:2d:2e:1e:2e:3b:f5:60_

![](http://j.ungeek.fr/eec74)

And to tell doenv to use it, I just have to type:
```
$ doenv keys + 1d:66:69:b5:d6:8e:0a:a6:4c:2d:2e:1e:2e:3b:f5:60
```

## Getting Started
Let's say I want to work on `MyFirtProject` using a 512mb instance running _Ruby on Rails on 14.04 (Postgres, Nginx, Unicorn)_.
First of all, I have to find the correct image to use.
```
$ doenv ls i
Fecthing images list...
Public images
    OS    |       SLUG       |              NAME              |    CREATION DATE     | MIN DISK
+---------+------------------+--------------------------------+----------------------+----------+
  CoreOS  | coreos-beta      | 835.5.0 (beta)                 | 2015-11-18T18:20:38Z |       20
  CentOS  | centos-5-8-x64   | 5.10 x64                       | 2014-09-26T16:40:18Z |       20
  CentOS  | centos-5-8-x32   | 5.10 x32                       | 2014-09-26T16:45:29Z |       20
  ...
  Ubuntu  | ruby-on-rails    | Ruby on Rails on 14.04         | 2015-09-03T15:02:16Z |       20
          |                  | (Postgres, Nginx, Unicorn)     |                      |       20
 ...
```
So the image's slug is _ruby-on-rails_. Then I create the instance.
```
$ doenv init MyFirstProject ruby-on-rails
Creating a droplet named MyFirstProject, Size: 512mb - Region: nyc2
... in-progress
... in-progress
... in-progress
... in-progress
... in-progress
... in-progress
... in-progress
... in-progress
... completed
Hoora! Just setup your droplet :)
Your IP is  162.243.112.188
```

And to start hacking on it, just run `doenv ssh MyFirstProject`. That's it. If you have configured your ssh-keys properly, they will be used. If that's not the case, you'll have to type the password sent to you by mail.

Two days pass, and you want to drop this and start working on something else ? Quite easy !
If you want to archive it,
```
$ doenv freeze MyFirstProject
Droplet powered on, stopping ...
... completed
Droplet is off :)
Creation of Snapshot with name MyFirstProject
... in-progress
...
... in-progress
... completed
Snapshot: done
Destroying droplet
Droplet destroyed ! :( That makes me sad ;(
```

And, if one day you want to work on it again ... `doenv restore MyFirstProject` ! :)

## Overview
```
$ doenv help
NAME:
   doenv - It's like virtualenv, but powered by DigitalOcean

USAGE:
   doenv [global options] command [command options] [arguments...]

VERSION:
   1.0.0

AUTHOR(S):
   PunKeel <punkeel@me.com>

COMMANDS:
   list, ls			List your droplets
   create, c, init, setup	Create a droplet (Initial setup)
   erase			Let's just drop this droplet.
   restore, bring, up		Create a droplet based on a snapshot (Restore snapshot, start)
   freeze			Freeze a droplet (Stop, Snapshot and Destroy)
   snapshot, snap, shot		Snapshot a droplet to reuse it later
   ssh, enter, hack		Hack into a droplet!
   keys				Manage the SSH keys used
   set-token			Set your Personnal API Token (PAT)
   help, h			Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

## Contribution Guidelines
Feel free to put up a pull request to fix a bug or maybe add a feature.
If I or any other collaborators agree that it is in line with the vision of the project, we will work with you to get
the code into a mergeable state and merge it into the master branch.

If you have contributed something significant to the project, I will most likely add you as a collaborator. As a
collaborator you are given the ability to merge others pull requests.

## Credits
- [DigitalOcean](https://digitalocean.com/) for making this possible,
- [digitalocean/godo](https://github.com/digitalocean/godo) for the easy to use Library,
- [codegangsta/cli](https://github.com/codegangsta/cli) for handling the CLI part of it,
- [Golang](https://golang.org/) for being cool and easy to learn.