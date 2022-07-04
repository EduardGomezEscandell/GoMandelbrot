"""
Prerequisites:
$ make
$ python -m pip install imageio
$ python -m pip install imageio-ffmpeg
"""

from dataclasses import dataclass
import subprocess
import math
import os
import shutil
import imageio.v2 as imageio

@dataclass(frozen=True)
class Settings:
    results_dir = 'results'
    resolution = (1088, 1088)
    centerpoint = "0+0i"
    colormap="pastel"
    batchsize = 8
    span = 4
    nframes = 200
    fps = 24
    format = "mp4"

def generate_command(settings):
    commands = [
        './mandelbrot -imw="{imw}" -imh="{imh}" -zc="{zc}" -zs="{zs}" -julia="{{real:.5f}}{{imag:+.5f}}i" -o "{resultsdir}/julia{{i}}.ppm" -c {cmap}',
        'pnmtopng "{resultsdir}/julia{{i}}.ppm" -quiet > "{resultsdir}/julia{{i}}.png"',
        'rm -f "{resultsdir}/julia{{i}}.ppm"',
        # 'echo "Done with frame {{i}}"'
    ]

    command = "{{{{ " + "; ".join(commands) + "; }}}}"
    return command.format(
        resultsdir=settings.results_dir,
        imw=settings.resolution[0],
        imh=settings.resolution[1],
        zc=settings.centerpoint,
        zs=settings.span,
        cmap=settings.colormap)

def startup(settings):
    try:
        shutil.rmtree(settings.results_dir)
    except FileNotFoundError:
        pass
    os.mkdir(settings.results_dir)


def generate_batch(command, begin, end, settings):
    processes = [None]*(end - begin)
    for i in range(begin, end):
        alpha = 2*math.pi*i/settings.nframes
        real = math.cos(alpha)
        imag = math.sin(alpha)
        cmd = command.format(real=real, imag=imag, i=i)
        processes[i - begin] = subprocess.Popen(cmd, shell = True)

    for p in processes:
        p.wait()

def animate(settings):
    with imageio.get_writer(f'{settings.results_dir}/movie.{settings.format}', mode='I', fps=settings.fps) as writer:
        for i in range(settings.nframes):
            image = imageio.imread(f"{settings.results_dir}/julia{i}.png")
            writer.append_data(image)
            os.remove(f"{settings.results_dir}/julia{i}.png")


def main():
    settings = Settings()
    command = generate_command(settings)
    startup(settings)

    print("Generating frames...")
    print('[', '-'*78, ']', sep='', end='\r')
    for i in range(0, settings.nframes, settings.batchsize):
        begin = i
        end = min(i+settings.batchsize, settings.nframes)
        generate_batch(command, begin, end, settings)
        prop = int(78*end/settings.nframes)
        print("[", '#' * prop, '-'*(78-prop),"]", sep='', end="\r")
    print()

    print("Assembling gif...")
    animate(settings)

    print("Done!")

if __name__ == '__main__':
    main()