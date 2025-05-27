# team-meter

***team-meter*** is a lightweight application designed to help manage and monitor DevOps team performance.

## Features

### 📊 Jira Integration
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

## Requirementes
- Docker
- Make
- Jira Credentials

## License
***team-meter*** released under MIT LICENSE.