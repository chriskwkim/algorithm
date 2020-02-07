# Tower of Hanoi

## Sample Run

$ ./tower-of-hanoi 4
Tower[0]: <4><3><2><1>
Tower[1]: 
Tower[2]: 
Move 4 disks from T[0] to T[2] using T[1]
Move 3 disks from T[0] to T[1] using T[2]
Move 2 disks from T[0] to T[2] using T[1]
Move 1 disks from T[0] to T[1] using T[2]
Moving Top disk 1 from T[0] to T[1]
Tower[0]: <4><3><2>
Tower[1]: <1>
Tower[2]: 
Moving Top disk 2 from T[0] to T[2]
Tower[0]: <4><3>
Tower[1]: <1>
Tower[2]: <2>
Move 1 disks from T[1] to T[2] using T[0]
Moving Top disk 1 from T[1] to T[2]
Tower[0]: <4><3>
Tower[1]: 
Tower[2]: <2><1>
Moving Top disk 3 from T[0] to T[1]
Tower[0]: <4>
Tower[1]: <3>
Tower[2]: <2><1>
Move 2 disks from T[2] to T[1] using T[0]
Move 1 disks from T[2] to T[0] using T[1]
Moving Top disk 1 from T[2] to T[0]
Tower[0]: <4><1>
Tower[1]: <3>
Tower[2]: <2>
Moving Top disk 2 from T[2] to T[1]
Tower[0]: <4><1>
Tower[1]: <3><2>
Tower[2]: 
Move 1 disks from T[0] to T[1] using T[2]
Moving Top disk 1 from T[0] to T[1]
Tower[0]: <4>
Tower[1]: <3><2><1>
Tower[2]: 
Moving Top disk 4 from T[0] to T[2]
Tower[0]: 
Tower[1]: <3><2><1>
Tower[2]: <4>
Move 3 disks from T[1] to T[2] using T[0]
Move 2 disks from T[1] to T[0] using T[2]
Move 1 disks from T[1] to T[2] using T[0]
Moving Top disk 1 from T[1] to T[2]
Tower[0]: 
Tower[1]: <3><2>
Tower[2]: <4><1>
Moving Top disk 2 from T[1] to T[0]
Tower[0]: <2>
Tower[1]: <3>
Tower[2]: <4><1>
Move 1 disks from T[2] to T[0] using T[1]
Moving Top disk 1 from T[2] to T[0]
Tower[0]: <2><1>
Tower[1]: <3>
Tower[2]: <4>
Moving Top disk 3 from T[1] to T[2]
Tower[0]: <2><1>
Tower[1]: 
Tower[2]: <4><3>
Move 2 disks from T[0] to T[2] using T[1]
Move 1 disks from T[0] to T[1] using T[2]
Moving Top disk 1 from T[0] to T[1]
Tower[0]: <2>
Tower[1]: <1>
Tower[2]: <4><3>
Moving Top disk 2 from T[0] to T[2]
Tower[0]: 
Tower[1]: <1>
Tower[2]: <4><3><2>
Move 1 disks from T[1] to T[2] using T[0]
Moving Top disk 1 from T[1] to T[2]
Tower[0]: 
Tower[1]: 
Tower[2]: <4><3><2><1>