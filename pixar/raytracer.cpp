#include <stdlib.h> // card > pixar.ppm
#include <stdio.h>
#include <math.h>
#define R return
#define O operator
typedef float F;typedef int I;struct V{F x,y,z;V(F v=0){x=y=z=v;}V(F a,F b,F
c=0){x=a;y=b;z=c;}V O+(V r){R V(x+r.x,y+r.y,z+r.z);}V O*(V r){R V(x*r.x,y*r.
y,z*r.z);}F O%(V r){R x*r.x+y*r.y+z*r.z;}V O!(){R*this*(1/sqrtf(*this%*this)
);}};F L(F l,F r){R l<r?l:r;}F U(){R(F)rand()/RAND_MAX;}F B(V p,V l,V h){l=p
+l*-1;h=h+p*-1;R-L(L(L(l.x,h.x),L(l.y,h.y)),L(l.z,h.z));}F S(V p,I&m){F d=1\
e9;V f=p;f.z=0;char l[]="5O5_5W9W5_9_COC_AOEOA_E_IOQ_I_QOUOY_Y_]OWW[WaOa_aW\
eWa_e_cWiO";for(I i=0;i<60;i+=4){V b=V(l[i]-79,l[i+1]-79)*.5,e=V(l[i+2]-79,l
[i+3]-79)*.5+b*-1,o=f+(b+e*L(-L((b+f*-1)%e/(e%e),0),1))*-1;d=L(d,o%o);}d=sq\
rtf(d);V a[]={V(-11,6),V(11,6)};for(I i=2;i--;){V o=f+a[i]*-1;d=L(d,o.x>0?f\
absf(sqrtf(o%o)-2):(o.y+=o.y>0?-2:2,sqrtf(o%o)));}d=powf(powf(d,8)+powf(p.z,
8),.125)-.5;m=1;F r=L(-L(B(p,V(-30,-.5,-30),V(30,18,30)),B(p,V(-25,17,-25),V
(25,20,25))),B(V(fmodf(fabsf(p.x),8),p.y,p.z),V(1.5,18.5,-25),V(6.5,20,25)))
;if(r<d)d=r,m=2;F s=19.9-p.y;if(s<d)d=s,m=3;R d;}I M(V o,V d,V&h,V&n){I m,s=
0;F t=0,c;for(;t<100;t+=c)if((c=S(h=o+d*t,m))<.01||++s>99)R n=!V(S(h+V(.01,0
),s)-c,S(h+V(0,.01),s)-c,S(h+V(0,0,.01),s)-c),m;R 0;}V T(V o,V d){V h,n,r,t=
1,l(!V(.6,.6,1));for(I b=3;b--;){I m=M(o,d,h,n);if(!m)break;if(m==1){d=d+n*(
n%d*-2);o=h+d*.1;t=t*.2;}if(m==2){F i=n%l,p=6.283185*U(),c=U(),s=sqrtf(1-c),
g=n.z<0?-1:1,u=-1/(g+n.z),v=n.x*n.y*u;d=V(v,g+n.y*n.y*u,-n.y)*(cosf(p)*s)+V(
1+g*n.x*n.x*u,g*v,-g*n.x)*(sinf(p)*s)+n*sqrtf(c);o=h+d*.1;t=t*.2;if(i>0&&M(h
+n*.1,l,h,n)==3)r=r+t*V(500,400,100)*i;}if(m==3){r=r+t*V(50,80,100);break;}}
R r;}I main(){I w=960,h=540,s=16;V e(-22,5,25),g=!(V(-3,4,0)+e*-1),l=!V(g.z,
0,-g.x)*(1./w),u(g.y*l.z-g.z*l.y,g.z*l.x-g.x*l.z,g.x*l.y-g.y*l.x);printf("P\
6 %d %d 255 ",w,h);for(I y=h;y--;)for(I x=w;x--;){V c;for(I p=s;p--;)c=c+T(e
,!(g+l*(x-w/2+U())+u*(y-h/2+U())));c=c*(1./s)+14./241;V o=c+1;c=V(c.x/o.x,c.
y/o.y,c.z/o.z)*255;printf("%c%c%c",(I)c.x,(I)c.y,(I)c.z);}}// Andrew Kensler
