# 🏦 Bond Search API (Mock)

A mock API built in Go to simulate bond search functionality. It allows filtering bonds by coupon rate and maturity date.

---

## 🔗 Endpoint

### GET `/api/bonds/search`

Query Parameters:

| Parameter        | Type   | Description                           |
|------------------|--------|---------------------------------------|
| `minCoupon`      | float  | Minimum coupon rate                   |
| `maxCoupon`      | float  | Maximum coupon rate                   |
| `maturityBefore` | date   | Filter bonds maturing **before** this date (YYYY-MM-DD) |
| `maturityAfter`  | date   | Filter bonds maturing **after** this date  (YYYY-MM-DD) |

---

## 🧪 Example Request

```http
GET /api/bonds/search?minCoupon=4.5&maturityBefore=2030-01-01
```

## Example response

```
[
  {
    "id": "BND002",
    "name": "Corporate Alpha",
    "issuer": "Ethio Corp",
    "couponRate": 4.8,
    "maturityDate": "2028-06-15",
    "price": 98.7
  },
  {
    "id": "BND004",
    "name": "Corporate Beta",
    "issuer": "Dashen Ltd",
    "couponRate": 5,
    "maturityDate": "2027-09-10",
    "price": 101.2
  }
]
```

## 🛠️ How to Run
Clone the repository:

```
git clone https://github.com/your-username/bond-api.git
cd bond-api
```

Run the server:

```
go run main.go
Visit: http://localhost:8080/api/bonds/search
```

## ✅ Acceptance Criteria
Filters return correct results

Works with any combination of parameters

Proper status codes and error messages on bad input

Code structure is clean and documented