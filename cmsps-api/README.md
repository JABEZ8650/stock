# Ethiopian CMSP Registry Mock API

This project simulates a CMSP registry using real Ethiopian CMSP data.  
Includes filtering by type and license date.

## Data Source

- [Ethio Telecom](https://www.ethiotelecom.et)
- [Safaricom Ethiopia](https://www.safaricom.et)

## API Endpoints:

- `GET /api/cmsps`
  - Optional query params: `type`, `licensedBefore`, `licensedAfter`
- `GET /api/cmsps/{id}`

## Run 

```bash
go mod tidy
go run main.go
```

