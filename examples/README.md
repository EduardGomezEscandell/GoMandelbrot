# Generation data

### Full
- Max iterations: 200
- Resolution: 1920 x 1080
- Subsampling level: 1
- Colormap: grayscale [0, 200]
- Complex plane center: -0.600000 +0.000000i
- Complex plane span:   4.000000 +2.250000i

![image](full.png)

### Spirals
- Max iterations: 200
- Resolution: 1920 x 1080
- Subsampling level: 1
- Colormap: grayscale [50, 200]
- Complex plane center: -0.803300 +0.178000i
- Complex plane span:   0.001700 +0.000956i

![image](spirals.png)

### Octopus
- Max iterations: 1500
- Resolution: 1920 x 1080
- Subsampling level: 1
- Colormap: grayscale [150, 1500] nl=100000
- Complex plane center: -0.802627 +0.1774515i
- Complex plane span:   2e-05 +1.125e-05i

![image](octopus.png)

### Octopus with sub-sampling
Same as before, but with:
- Subsampling level: 5

Acts as an antialiasing. Performance impact is O(n²), with n being the subsampling level.

![image](octopus_ss.png)

### Copies
- Max iterations: 200
- Resolution: 1920 x 1080
- Subsampling level: 1
- Colormap: grayscale [0, 200]
- Complex plane center: -0.900000 +0.250000i
- Complex plane span:   0.187500 +0.105469i

![image](copies.png)

### Copies (bright repetitive colormap)
- Max iterations: 50
- Resolution: 1920 x 1080
- Subsampling level: 3
- Colormap: multicolor [0, 50] nl=0
- Complex plane center: -0.9+0.25i
- Complex plane span:   0.1875+0.10546875i

![image](copies_multicolor.png)

### Copies (pastel repetitive colormap)
- Max iterations: 50
- Resolution: 1920 x 1080
- Subsampling level: 3
- Colormap: pastel [0, 50] nl=0
- Complex plane center: -0.9+0.25i
- Complex plane span:   0.1875+0.10546875i

![image](copies_pastel.png)


## Mini
- Max iterations: 400
- Resolution: 1920 x 1080
- Subsampling level: 1
- Colormap: grayscale [0, 400]
- Complex plane center: -1.256000 +0.382000i
- Complex plane span:   0.015000 +0.008438i

![image](mini.png)

### Maze
- Max iterations: 5000
- Resolution: 1920 x 1080
- Subsampling level: 3
- Colormap: grayscale [400, 5000] nl=100000
- Complex plane center: 0.356+0.331i
- Complex plane span:   0.0001+5.625e-05i

![image](maze.png)
