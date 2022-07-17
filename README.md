![image](examples/banner.png)

This is a practise exercise to learn Go. It draws the mandelbrot and Julia sets. You can read the theory and details, or you can skip directly to the images [at the end of this page](#figures).

## The math
### Mandelbrot set
The [mandelbrot set](https://en.wikipedia.org/wiki/Mandelbrot_set) *M ⊂ ℂ*, is a mathematical object defined as the set of complex points *c* such that:
- Starting with *z = 0*
- Repeteadly applying the transformation *T: z → z² + c*, the norm does not diverge to infinty.

It is known that if *|z| > 2*, then the value diverges, hence we can test whether a point belongs to the set by checking if *|T(T(T(...T(z)...)))| ≤ 2*. To be certain that a particular point does not diverge, we'd have to iterate at infinitum, hence an upper bound must be set. This upper bound controls the accuracy (higher is better) and computational expense (lower is cheaper).

### Julia set
The [Julia set](https://en.wikipedia.org/wiki/Julia_set#Quadratic_polynomials) *J(f) ⊂ ℂ*, is defined with respect to a function *f: ℂ → ℂ*. This codebase particularizes to *f(z)* being a quadratic polynomial with the form *p(z) = z² + k*. The Julia set of a function *f* is defined as the set of all points *c ∈ ℂ* such that:
- Starting with *z = c*
- Repeteadly applying the transformation *T: z → f(z)*, the norm does not diverge to infinty.

We can investigate related Julia sets by varying the parameter *k* of the quadratic polynomial. This is what `scripts/julia_animation.py` does, and is shown in the video below.

### Colormaps
We can define a function *D: ℂ → ℕ*, that maps each point in the complex plane to the number of iterations it takes to reach *|z|≥ 2*. The pictures generated are representations of this field. To convert them to colours, we can take two approaches.
1. Monotonous colormaps: A colour gradient varies with *D(z)*. Lighter colors represent faster divergence times (small *D(z)*), and pure black represents that the maximum number of iterations has been reached.
2. Modular colormaps: Given a set of *M* colours, plus black:
    - The black region is the approximation of the Mandelbrot or Julia set (i.e. *D(z) >= maxiter*)
    - Each colour other than black represents a value of *D(z) mod M*.

Monotonous colormaps (such as the grayscale one here) are better at showing divergence period, whereas modular or repetitive colourmaps are better at showing the borders between regions.

## How-to

### How to run
The first time you must build it:
```
make
```
Then, to run it, you do
```
./mandelbrot
```
You can use flag `--help` to see the tweakables.


### Image format
The image is outputted as a [NetPBM](https://en.wikipedia.org/wiki/Netpbm) image, with either format
- `.bin.ppm` OR `.ppm` (default)
- `.ascii.ppm`

You can choose the format with `-o FILENAME.FORMAT`, but the only reason to use ASCII is to debug.
You can convert them to png with:
```bash
pnmtopng FILENAME.FORMAT > FILENAME.png
```

## Figures
Here are a few images generated by this project. Go to [examples](examples/README.md) to see more examples, and technical details about them.

![image](examples/full.png)
Mandelbrot set with a monotonous colormap.

![image](examples/julia.gif)
Exploring the Julia set with a modular colormap.

![image](examples/spirals_pastel.png)
An interesting region of the Mandelbrot set, with a modular colormap.

![image](examples/octopus.png)
An interesting region of the Mandelbrot set, with a monotonous colormap.

![image](examples/copies_multicolor.png)
An interesting region of the Mandelbrot set, with a modular colormap.

![image](examples/mini.png)
An region of the Mandelbrot showing self-similarity, with a monotonous colormap.
