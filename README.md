# arch-decision-records

CLI tool to manage ADRs.

## What is an ADR?

An Architectural Decision Record is a short, well defined and structured document describing the process leading to any significant architectural decision. You can read more about it [here](http://thinkrelevance.com/blog/2011/11/15/documenting-architecture-decisions).

## Usage

Add a new directory to your team with a `template.md`, or run:

```
adr create-team TEAM
```

Keep the team code up to a few uppercase characters, it will be used by other people referring to your ADR.

Copy `template.md` to `TEAM/0002-title-separated-by-dashes.md`, keeping this very same format, or run:

```
adr create --team TEAM "Title as you would normally write"
```

## ADR tool installation

Download the latest release [here](https://github.com/inloco/adr-tool/releases) to install `adr`.

Copy the following commands to your terminal to make it executable:

**Mac**
```
tar -xzvf ~/Downloads/adr-tool_0.0.2_darwin_amd64.tar.gz
mv adr-tool /usr/local/bin/adr
chmod +x /usr/local/bin/adr
```

**Linux**
```
tar -xzvf ~/Downloads/adr-tool_0.0.2_linux_amd64.tar.gz
mv adr-tool /usr/local/bin/adr
chmod +x /usr/local/bin/adr
```
