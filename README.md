# team-meter

***team-meter*** is a lightweight application designed to help manage and monitor DevOps team performance.

## Features

### Jira Integration
- Process and analyze Jira issues.
- Visual dashboards with key delivery metrics:
  - **Cycle Time**
  - **Lead Time**
  - **Work in Progress (WIP)**
  - **Throughput**

## Getting Started

1. Clone the repository:
```bash
  $ git clone https://github.com/ambrisolla/team-meter.git
  $ cd team-meter
```
2. Configure .env file with your own information
```bash
  $ cp .env.example .env
```
3. Start the stack
```bash
  $ make up
```
4. Stop the stack and remove Docker resources
```bash
  $ make down
```

## Access dashboard
After starting the stack, access Grafana in your browser:
- ```http://localhost:3000```

## Requirementes
- Docker
- Make
- Jira Credentials

## Throubleshootings
If Grafana does not display any data, follow the steps below:

- Check if all containers are up:
```bash
$ docker ps                                                                                                                                                                                                    
CONTAINER ID   IMAGE             COMMAND                  CREATED          STATUS          PORTS                    NAMES
db89cda23d43   team-meter-app    "/app/app"               20 minutes ago   Up 20 minutes                            team_meter_app
1a9cec58cc94   postgres:17.5     "docker-entrypoint.s…"   20 minutes ago   Up 20 minutes   0.0.0.0:5432->5432/tcp   team_meter_db
f19dca4645d5   grafana/grafana   "/run.sh"                20 minutes ago   Up 20 minutes   0.0.0.0:3000->3000/tcp   team_meter_grafana
```
- Check for error in **team_meter_app**  container logs:
```bash
$ docker logs -f team_meter_app                                                                                                                                                                                                                                                                                                                                                                                           
{"level":"info","ts":1748382611.2449064,"caller":"jira-issues/main.go:21","msg":"starting application","service":"team-meter","version":"0.0.1"}
{"level":"info","ts":1748382611.2626405,"caller":"jira/issues.go:23","msg":"starting processing FPDEE issues","service":"team-meter","version":"0.0.1"}
{"level":"info","ts":1748382619.635832,"caller":"jira/issues.go:38","msg":"Processing 100 issues for project FPDEE (page 1)","service":"team-meter","version":"0.0.1"}
{"level":"info","ts":1748382628.3026667,"caller":"jira/issues.go:38","msg":"Processing 100 issues for project FPDEE (page 2)","service":"team-meter","version":"0.0.1"}
{"level":"info","ts":1748382637.9074538,"caller":"jira/issues.go:38","msg":"Processing 100 issues for project FPDEE (page 3)","service":"team-meter","version":"0.0.1"}
{"level":"info","ts":1748382646.9587367,"caller":"jira/issues.go:38","msg":"Processing 100 issues for project FPDEE (page 4)","service":"team-meter","version":"0.0.1"}
{"level":"info","ts":1748382655.9250968,"caller":"jira/issues.go:38","msg":"Processing 100 issues for project FPDEE (page 5)","service":"team-meter","version":"0.0.1"}
{"level":"info","ts":1748382664.9650438,"caller":"jira/issues.go:38","msg":"Processing 100 issues for project FPDEE (page 6)","service":"team-meter","version":"0.0.1"}
{"level":"info","ts":1748382674.8304682,"caller":"jira/issues.go:38","msg":"Processing 100 issues for project FPDEE (page 7)","service":"team-meter","version":"0.0.1"}
```

## License
***team-meter*** released under MIT LICENSE.