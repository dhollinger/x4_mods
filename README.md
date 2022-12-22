# Changes made for VRO

## Ship Names

A few of the ship names are already used in VRO:

* VRO contains an 'L' Frigate ported from X3 called the Centaur.
* VRO already contains a Split Battleship called the Gharial.
* VRO already contains a Teladi Battleship called the Fulmar.

As such the following changes were made:

* The Centaur battleship is now the Aamon Battleship
* The Gharial battleship (this mod) is now the Copperhead Battleship (Sticking with the reptile/snake theme).
* The Fulmar battleship is now called the Raven Battleship (sticking with the avian theme).

## Ship Changes

### Battleships

The battleships in this mod are much "heavier" in terms of their sheer firepower than the base VRO battleships. As such, the following balance changes were made:

* The hull of each battleship has been increased and they are now tougher than the base VRO battleships.
* The mass and drag of each battleship has been increased.
* In some cases, missile count was increased and unit count decreased.
* Radar range of each battleship in this mod has been increased to greater than that of the battleships in base VRO.
* Min, Avg, and Max prices have been updated relative to VRO base prices for other Capital Ships.

### Heavy Frigates

The Frigates in this mod add a new type of Frigate to VRO over the more drone/escort focused Frigates in base VRO. As such, they are heavier, tougher, and a bit slower than their VRO counterparts, but pack more punch.

* The hull, mass, and drag have been increased over their VRO Escort Frigate counterparts.
* The M weapons on the Caiman have been given the `hgun` tags to ensure that it can make use of the VRO Heavy Medium Weapons. This seemed in-line with rest of the ship's capabilities.
* Min, Avg, Max prices updated relative to the VRO base prices for other Frigates.

### Gunboats/Blockade Runner

Again, the Gunboats in this mod fill an empty gap where heavier gunboats don't yet exist. As such, their stats make them slightly slower and more expensive than the base VRO counterparts.

The Baku blockade runner was left largely alone, aside from increasing the base hull to match its role relative to other ship hulls in VRO. This ship is slightly weaker, but faster than other ships of its type.

* All M sized weapons have been given the `standard` and `highpoer` tags to take full advantage of all the M sized guns in VRO.
* The Heavy Gunboats have split the difference between the manueverability and speed of vanilla frigates/gunboats and corvettes. Making them slightly more manueverable than the former, but less than the latter.
* Ensured the Blockade runner's hull, mass, and drag were within a range that ensure they keep their speed and manueverability in VRO.

## Building

From the project directory execute the following command in Linux:

```bash
$ tools/build.sh
```

A `.bat` file will be added soon