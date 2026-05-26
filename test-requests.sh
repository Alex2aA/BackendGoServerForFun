#!/bin/bash

BASE_USER="http://localhost:8081"
BASE_BOOKING="http://localhost:8082"

echo "=== 1. Register user ==="
curl -s -X POST "$BASE_USER/api/user/register" \
  -H "Content-Type: application/json" \
  -d '{"username":"john","password":"secret123"}' | python3 -m json.tool

echo ""
echo "=== 2. Login user ==="
curl -s -X POST "$BASE_USER/api/user/login" \
  -H "Content-Type: application/json" \
  -d '{"username":"john","password":"secret123"}' | python3 -m json.tool

echo ""
echo "=== 3. Create hostel ==="
curl -s -X POST "$BASE_BOOKING/api/hostel" \
  -H "Content-Type: application/json" \
  -d '{"name":"Sunset Hostel","description":"Cozy hostel near the beach","rate":4}' | python3 -m json.tool

echo ""
echo "=== 4. Create house ==="
curl -s -X POST "$BASE_BOOKING/api/house" \
  -H "Content-Type: application/json" \
  -d '{
    "hostel_id": "hostel-001",
    "address": "123 Beach St",
    "number_of_rooms": 3,
    "price_per_day": 50,
    "name": "Beach House",
    "description": "Nice house with ocean view"
  }' | python3 -m json.tool

echo ""
echo "=== 5. Create booking ==="
curl -s -X POST "$BASE_BOOKING/api/booking" \
  -H "Content-Type: application/json" \
  -d '{
    "house_id": "house-001",
    "user_id": "user-001",
    "date_start": "2026-06-01",
    "date_end": "2026-06-05"
  }' | python3 -m json.tool

echo ""
echo "=== 6. Swagger user-service ==="
curl -s -o /dev/null -w "HTTP %{http_code} -> %{redirect_url}" "$BASE_USER/swagger/"

echo ""
echo "=== 7. Swagger booking-service ==="
curl -s -o /dev/null -w "HTTP %{http_code} -> %{redirect_url}" "$BASE_BOOKING/swagger/"
echo ""
