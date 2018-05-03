#!/bin/bash

storeid="$1"

while IFS='' read -r line || [[ -n "$line" ]]; do
  id=$(printf "$line" | cut -d',' -f1 | xargs)
  name=$(printf "$line" | cut -d',' -f2 | xargs)
  size=$(printf "$line" | cut -d',' -f3 | xargs)

  printf "
    INSERT INTO starbucks.products
      (id, name, size)
    values
      ($id, '$name', '$size');
    INSERT INTO starbucks.inventory
      (store_id, id, name, size, quantity)
    values
      ($storeid, $id, '$name', '$size', '10000');" | cqlsh
done < "$2"
