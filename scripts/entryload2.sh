#!/usr/bin/env bash

for j in `seq 1 2800`;
do
 echo $loopi $j "entry" | factom-cli2 put -c ${chain1} e1 &
 echo $loopi $j "entry" | factom-cli2 put -c ${chain2} e1 &
 echo $loopi $j "entry" | factom-cli2 put -c ${chain3} e1 &
 sleep .02s
done
echo $loopi complete!

