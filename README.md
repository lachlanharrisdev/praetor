<!--<br />-->
<div align="center">
  <!--
  <a href="https://github.com/lachlanharrisdev/praetor">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>
  -->

<h1 align="center">Praetor</h1>

  <p align="center">
    Sophisticated engagement management & low-friction notetaking for penetration testing
    <br />
    <a href="https://github.com/lachlanharrisdev/praetor"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/lachlanharrisdev/praetor/blob/main/CONTRIBUTING.md">Contribute</a>
    &middot;
    <a href="https://github.com/lachlanharrisdev/praetor/issues/new?template=bug_report.md">Report Bug</a>
    &middot;
    <a href="https://github.com/lachlanharrisdev/praetor/issues/new?template=feature_request.md">Request Feature</a>
  </p>
</div>

<br/>
<h2>What is Praetor?</h2>

Praetor is a Go-based CLI tool built solely to reduce administrative friction in penetration testing. `pt` offers a clean set of integrated utilities to:
* Manage engagement contexts & directories
* Intuitive commands to record, update, report on, format and export commands, notes & outputs
* A lightweight, centralised & forward-immutable event log to contain context for each engagement
* And endless possibilities to customise to suit Praetor to your team's needs and reduce the cognitive load of administrative work.

<br/>
<h2 align="right">How do I get Started?</h2>

<h3>Installation</h3>

1. Go to the [releases](https://github.com/lachlanharrisdev/praetor/releases/) page and download the desired version & `checksums.txt` file. It should look like `praetor_{version}_{os}_{arch}.targ.gz`
2. Run the following commands in your shell to extract and move to your `bin`
```sh
tar xzf praetor_0.0.4-dev.1_<os>_<arch>.tar.gz
sudo mv pt /usr/local/bin/
```
3. Verify the checksums
```sh
sha256sum -c checksums.txt
```
4. Verify the installation succeeded
```sh
pt version
```

<br/>
<h3>Usage</h3>

`pt` has countless methods of use. There's no one correct way to use it and it all depends on your environment, existing methods of administration and the needs of your team. It's best to keep up with the documentation and refer to each individual commands use.

Some basic usage could look as follows:

1. Create a new engagement directory
```bash
$ pt init test-eng
/home/{user}/engagements/test-eng/

$ cd test-eng
```
2. Take your first note
```bash
$ pt note Engagement begun. Provided IP: 123.45.67.89
```
3. Record a tool output
```bash
$ nmap -sC 123.45.67.89 | pt capture

# or:
$ nmap -sC -o nmap_result.txt 123.45.67.89
$ pt capture nmap_result.txt

# or:
$ pt run nmap -sC 123.45.67.89
```
4. View the last few events
```bash
$ pt list 3
```

---

<br/>

> This project is in a WIP state. It is currently unstable and not recommended for use within automated systems or under strict compliance policies. All code is open source and aims to have a minimal, secure footprint, but in it's pre-release stages no guarantees can be made.
>
> This project is licensed under the GPL-3.0 License. Please see [LICENSE](https://github.com/lachlanharrisdev/praetor?tab=GPL-3.0-1-ov-file) for more info.
>
> Copyright © Lachlan Harris 2026. All Rights Reserved.
