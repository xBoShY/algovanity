AlgoVanity generates an Algorand Account where the prefix for the account address is chosen by the user.

## Build with docker
```bash
$ docker build . -t xboshy/algovanity
```

## Run docker container
```bash
$ docker run --rm -e PREFIX=ANGRA -e THREADS=28 -e COUNT=4 xboshy/algovanity
```
Result:
```bash
PREFIX: ANGRA
THREADS: 28
COUNT: 2
====================
Generations: 97577514
Time: 3m13.852796911s
Address: ANGRARIUEXTLO2AGUM3AM5ULWBKPNRABECHD5TPDKA63P2HAFTMWGYPSJA
Mnemonic: young siren creek ugly fat rule lamp orange crater endorse ship cream exit bar stock source depart side have mind innocent tent captain absent old
PublicKey: [3 77 16 69 20 37 230 183 104 6 163 54 6 118 139 176 84 246 196 1 32 142 62 205 227 80 61 183 232 224 44 217]
PrivateKey: [250 95 114 102 192 190 169 244 154 207 155 148 113 18 140 47 243 167 73 192 26 208 214 249 113 211 206 72 58 125 71 164 3 77 16 69 20 37 230 183 104 6 163 54 6 118 139 176 84 246 196 1 32 142 62 205 227 80 61 183 232 224 44 217]
====================
Generations: 108295192
Time: 3m35.176847643s
Address: ANGRARB2I3Y35QN2MZ2T7HPMSP6CE42PUEZTKTAWSRHAVBJIRYWZ7J2EYQ
Mnemonic: web code rack salmon please tail bird december pet typical desk such sea raw friend glad skirt sugar anxiety square unit employ twin able obvious
PublicKey: [3 77 16 68 58 70 241 190 193 186 102 117 63 157 236 147 252 34 115 79 161 51 53 76 22 148 78 10 133 40 142 45]
PrivateKey: [198 71 11 97 231 59 83 117 211 162 56 27 253 250 119 134 253 96 202 158 139 98 85 54 54 20 54 221 246 35 109 93 3 77 16 68 58 70 241 190 193 186 102 117 63 157 236 147 252 34 115 79 161 51 53 76 22 148 78 10 133 40 142 45]
```
