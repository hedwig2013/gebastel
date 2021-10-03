from PIL import Image as image
from PIL import ImageDraw as draw
import os
from datetime import datetime
from datetime import timedelta
import sys
import json

dur1=timedelta()
dur2=timedelta()

report=list()

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
            s2=datetime.now()
            width=i.width
            height=i.height
            d=draw.Draw(i)
            d.line((0,0,width,height),fill=(255,255,255),width=5)
            d.line((width,0,0,height),fill=(255,255,255),width=5)
            s3=datetime.now()

            # i.save("review.png",format="png")

            dur1+=(s2-s1)
            dur2+=(s3-s2)
            statresult=os.stat(pl)
            bs=statresult.st_size

            rp=dict()
            rp["extension"]=ext
            rp["bytesize"]=bs
            rp["loadtime"]=(s2-s1).microseconds
            rp["drawtime"]=(s3-s2).microseconds
            report.append(rp)
        except Exception:
            continue
        finally:
            i.close()

rs=json.dumps(report)

print(rs)
print(dur1)
print(dur2)
