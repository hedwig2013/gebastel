from PIL import Image as image
from PIL import ImageDraw as draw
import os
from datetime import datetime
from datetime import timedelta
import sys

dur=timedelta()

if len(sys.argv)!=2:
    print("Bitte genau einen Pfad zu einer Bildersammlung angeben.")
    sys.exit()

b=sys.argv[1]

for r,d,f in os.walk(b):
    for file in f:
        ext=os.path.splitext(file)[1]

        pl=os.path.join(r,file)

        try:
            s1=datetime.now()
            i=image.open(pl)
            width=i.width
            height=i.height
            d=draw.Draw(i)
            d.line((0,0,width,height),fill=(255,255,255),width=5)
            d.line((width,0,0,height),fill=(255,255,255),width=5)
            s2=datetime.now()

            i.save("review.png",format="png")

            dur+=(s2-s1)
        except Exception:
            continue
        finally:
            i.close()

print(dur)
