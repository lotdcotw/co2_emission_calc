echo ./bin/client -start Hamburg -end Berlin -transportation-method medium-diesel-car
./bin/client -start Hamburg -end Berlin -transportation-method medium-diesel-car

echo ./bin/client --start Hamburg --end Berlin --transportation-method medium-diesel-car
./bin/client --start Hamburg --end Berlin --transportation-method medium-diesel-car

echo ./bin/client --end Berlin --transportation-method medium-diesel-car --start Hamburg
./bin/client --end Berlin --transportation-method medium-diesel-car --start Hamburg

echo ./bin/client --end Berlin    --transportation-method medium-diesel-car       --start Hamburg
./bin/client --end Berlin    --transportation-method medium-diesel-car       --start Hamburg

echo ./bin/client --start=Hamburg --end=Berlin --transportation-method=medium-diesel-car
./bin/client --start=Hamburg --end=Berlin --transportation-method=medium-diesel-car

echo ./bin/client --start="New York" --end="Los Angeles" --transportation-method=medium-diesel-car
./bin/client --start="New York" --end="Los Angeles" --transportation-method=medium-diesel-car

echo ./bin/client --start "New York" --end "Los Angeles" --transportation-method=medium-diesel-car
./bin/client --start "New York" --end "Los Angeles" --transportation-method=medium-diesel-car