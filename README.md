# gebastel
Mal einfach rumexperimentieren.

NOTHING TO EXPLAIN. IT IS A PRIVATE PLAYGROUND!

Python:

Bildladen:      0m3.144971s

Kreuz zeichnen: 1m32.665759s

Go mit std pkg image:

Bildladen:      3m51.3862602s

Kreuz zeichnen: 1m12.3785268s

Hier nun Go mit gocv und OpenCV 4.5.3:

Bildladen:      2m32.0122601s

Kreuz zeichnen: 0m7.8792553s

Gleiche Bildsammlung. Python ist immer noch schneller. WTF?

    ----------------------------------------------------------+-------------
        flat  flat%   sum%        cum   cum%   calls calls% + context
    ----------------------------------------------------------+-------------
                                                98.42s 94.81% |   gocv.io/x/gocv._Cfunc_Image_IMRead
                                                4.07s  3.92% |   gocv.io/x/gocv._Cfunc_Line
                                                0.82s  0.79% |   gocv.io/x/gocv._Cfunc_Mat_Close
                                                0.37s  0.36% |   main.walker
                                                0.11s  0.11% |   main.main
    103.78s 96.58% 96.58%    103.81s 96.60%                | runtime.cgocall
    ----------------------------------------------------------+-------------
                                                1.41s   100% |   runtime.sysUnused
        1.41s  1.31% 97.89%      1.41s  1.31%                | runtime.stdcall3
    ----------------------------------------------------------+-------------
                                                0.79s 98.75% |   runtime.park_m
        0.07s 0.065% 97.95%      0.80s  0.74%                | runtime.schedule
    ----------------------------------------------------------+-------------
                                                0.64s   100% |   runtime.resettimer
        0.02s 0.019% 97.97%      0.64s   0.6%                | runtime.modtimer
                                                0.57s 89.06% |   runtime.wakeNetPoller
    ----------------------------------------------------------+-------------
                                                4.08s   100% |   gocv.io/x/gocv.Line.func1
        0.01s 0.0093% 97.98%      4.08s  3.80%                | gocv.io/x/gocv._Cfunc_Line
                                                4.07s 99.75% |   runtime.cgocall
    ----------------------------------------------------------+-------------
                                                1.44s   100% |   runtime.(*pageAlloc).scavenge
        0.01s 0.0093% 97.99%      1.44s  1.34%                | runtime.(*pageAlloc).scavengeOne
                                                1.42s 98.61% |   runtime.(*pageAlloc).scavengeRangeLocked
    ----------------------------------------------------------+-------------
                                                0.97s   100% |   runtime.mcall
        0.01s 0.0093% 98.00%      0.97s   0.9%                | runtime.park_m
                                                0.79s 81.44% |   runtime.schedule
                                                0.14s 14.43% |   runtime.resettimer (inline)
    ----------------------------------------------------------+-------------
                                                1.42s   100% |   runtime.(*pageAlloc).scavengeRangeLocked
        0.01s 0.0093% 98.01%      1.42s  1.32%                | runtime.sysUnused
                                                1.41s 99.30% |   runtime.stdcall3
    ----------------------------------------------------------+-------------
                                                0.83s   100% |   main.walker
            0     0% 98.01%      0.83s  0.77%                | gocv.io/x/gocv.(*Mat).Close
                                                0.83s   100% |   gocv.io/x/gocv.(*Mat).Close.func1
    ----------------------------------------------------------+-------------
                                                0.83s   100% |   gocv.io/x/gocv.(*Mat).Close
            0     0% 98.01%      0.83s  0.77%                | gocv.io/x/gocv.(*Mat).Close.func1
                                                0.82s 98.80% |   gocv.io/x/gocv._Cfunc_Mat_Close
    ----------------------------------------------------------+-------------
                                                98.45s   100% |   main.walker
            0     0% 98.01%     98.45s 91.62%                | gocv.io/x/gocv.IMRead
                                                98.42s   100% |   gocv.io/x/gocv._Cfunc_Image_IMRead
    ----------------------------------------------------------+-------------
                                                4.08s   100% |   main.walker
            0     0% 98.01%      4.08s  3.80%                | gocv.io/x/gocv.Line
                                                4.08s   100% |   gocv.io/x/gocv.Line.func1
    ----------------------------------------------------------+-------------
                                                4.08s   100% |   gocv.io/x/gocv.Line
            0     0% 98.01%      4.08s  3.80%                | gocv.io/x/gocv.Line.func1
                                                4.08s   100% |   gocv.io/x/gocv._Cfunc_Line
    ----------------------------------------------------------+-------------
                                                98.42s   100% |   gocv.io/x/gocv.IMRead
            0     0% 98.01%     98.42s 91.59%                | gocv.io/x/gocv._Cfunc_Image_IMRead
                                                98.42s   100% |   runtime.cgocall
    ----------------------------------------------------------+-------------
                                                0.82s   100% |   gocv.io/x/gocv.(*Mat).Close.func1
            0     0% 98.01%      0.82s  0.76%                | gocv.io/x/gocv._Cfunc_Mat_Close
                                                0.82s   100% |   runtime.cgocall
    ----------------------------------------------------------+-------------
                                            103.88s   100% |   runtime.main
            0     0% 98.01%    103.88s 96.67%                | main.main
                                            103.76s 99.88% |   main.walker
                                                0.11s  0.11% |   runtime.cgocall
    ----------------------------------------------------------+-------------
                                            103.76s   100% |   main.main
            0     0% 98.01%    103.76s 96.56%                | main.walker
                                                98.45s 94.88% |   gocv.io/x/gocv.IMRead
                                                4.08s  3.93% |   gocv.io/x/gocv.Line
                                                0.83s   0.8% |   gocv.io/x/gocv.(*Mat).Close
                                                0.37s  0.36% |   runtime.cgocall
    ----------------------------------------------------------+-------------
                                                1.45s   100% |   runtime.bgscavenge.func2
            0     0% 98.01%      1.45s  1.35%                | runtime.(*pageAlloc).scavenge
                                                1.44s 99.31% |   runtime.(*pageAlloc).scavengeOne
    ----------------------------------------------------------+-------------
                                                1.42s   100% |   runtime.(*pageAlloc).scavengeOne
            0     0% 98.01%      1.42s  1.32%                | runtime.(*pageAlloc).scavengeRangeLocked
                                                1.42s   100% |   runtime.sysUnused
    ----------------------------------------------------------+-------------
            0     0% 98.01%      0.55s  0.51%                | runtime.bgscavenge
                                                0.50s 90.91% |   runtime.resettimer (inline)
    ----------------------------------------------------------+-------------
                                                1.45s   100% |   runtime.systemstack
            0     0% 98.01%      1.45s  1.35%                | runtime.bgscavenge.func2
                                                1.45s   100% |   runtime.(*pageAlloc).scavenge
    ----------------------------------------------------------+-------------
            0     0% 98.01%    103.88s 96.67%                | runtime.main
                                            103.88s   100% |   main.main
    ----------------------------------------------------------+-------------
            0     0% 98.01%      0.97s   0.9%                | runtime.mcall
                                                0.97s   100% |   runtime.park_m
    ----------------------------------------------------------+-------------
                                                0.50s 78.12% |   runtime.bgscavenge (inline)
                                                0.14s 21.88% |   runtime.park_m (inline)
            0     0% 98.01%      0.64s   0.6%                | runtime.resettimer
                                                0.64s   100% |   runtime.modtimer
    ----------------------------------------------------------+-------------
            0     0% 98.01%      1.49s  1.39%                | runtime.systemstack
                                                1.45s 97.32% |   runtime.bgscavenge.func2
    ----------------------------------------------------------+-------------
                                                0.57s   100% |   runtime.modtimer
            0     0% 98.01%      0.57s  0.53%                | runtime.wakeNetPoller
    ----------------------------------------------------------+-------------
            0     0% 98.01%      0.59s  0.55%                | runtime/pprof.profileWriter
    ----------------------------------------------------------+-------------
