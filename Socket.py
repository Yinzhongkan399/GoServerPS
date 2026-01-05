import sqlite3
import os
import time
import json
import sys

def TranV4intoP(inputStr:str):
    r=inputStr.split(":")
    firstByte=int(r[0][6:8],16)
    secByte=int(r[0][4:6],16)
    thridByte=int(r[0][2:4],16)
    lastByte=int(r[0][0:2],16)
    port=int(r[1],16)
    return str(firstByte)+"."+str(secByte)+"."+str(thridByte)+"."+str(lastByte)+":"+str(port)

def TranV6intoP(inputStr:str):
    r=inputStr.split(":")
    port=int(r[1],16)
    return r[0][0:4].lower()+":"+r[0][4:8].lower()+":"+r[0][8:12].lower()+":"+r[0][12:16].lower()+\
        ":"+r[0][16:20].lower()+":"+r[0][20:24].lower()+":"+r[0][24:28].lower()+":"+r[0][28:32].lower()+":"+str(port)

def TranStateintoSTR(inputint:int):
    if inputint==1:
        return "01(ESTABLISHED)"
    elif inputint==2:
        return "02(SYN_SENT)"
    elif inputint==3:
        return "03(SYN_RECV)"
    elif inputint==4:
        return "04(FIN_WAIT1)"
    elif inputint==5:
        return "05(FIN_WAIT2)"
    elif inputint==6:
        return "06(TIME_WAIT)"
    elif inputint==7:
        return "07(CLOSE)"
    elif inputint==8:
        return "08(CLOSE_WAIT)"
    elif inputint==9:
        return "09(LAST_ACK)"
    elif inputint==10:
        return "0A(LISTEN)"
    elif inputint==11:
        return "0B(CLOSING)"
    return str(inputint)+("(UNDEFINED)")
    

TotalData={}
def GetTcpInfo():
    f=open("/proc/net/tcp","r")
    data=[]
    for item in f.readlines()[1:]:
        r=item.split()
        data.append((curTime,r[0][:-1],TranV4intoP((r[1])),TranV4intoP((r[2])),TranStateintoSTR(int(r[3],16))))
    TotalData["tcpipv4"]=data
    f.close()
    data=[]
    f=open("/proc/net/tcp6","r")
    for item in f.readlines()[1:]:
        r=item.split()
        data.append((curTime,r[0][:-1],TranV6intoP(r[1]),TranV6intoP(r[2]),TranStateintoSTR(int(r[3],16))))
    TotalData["tcpipv6"]=data
    f.close()
    pass

def GetUdpInfo():
    f=open("/proc/net/udp","r")
    data=[]
    for item in f.readlines()[1:]:
        r=item.split()
        data.append((curTime,r[0][:-1],TranV4intoP((r[1])),TranV4intoP((r[2])),TranStateintoSTR(int(r[3],16))))
    TotalData["udpipv4"]=data
    f.close()
    f=open("/proc/net/udp6","r")
    data=[]
    for item in f.readlines()[1:]:
        r=item.split()
        data.append((curTime,r[0][:-1],TranV6intoP(r[1]),TranV6intoP(r[2]),TranStateintoSTR(int(r[3],16))))
    TotalData["udpipv6"]=data
    f.close()
    pass

def GetIcmpInfo():
    f=open("/proc/net/icmp","r")
    data=[]
    for item in f.readlines()[1:]:
        r=item.split()
        data.append((curTime,r[0][:-1],TranV4intoP((r[1])),TranV4intoP((r[2])),TranStateintoSTR(int(r[3],16))))
    f.close()
    TotalData["icmpipv4"]=data
    f=open("/proc/net/icmp6","r")
    data=[]
    for item in f.readlines()[1:]:
        r=item.split()
        data.append((curTime,r[0][:-1],TranV6intoP(r[1]),TranV6intoP(r[2]),TranStateintoSTR(int(r[3],16))))
    TotalData["icmpipv6"]=data
    f.close()
    pass

def GetRawInfo():
    f=open("/proc/net/raw","r")
    data=[]
    for item in f.readlines()[1:]:
        r=item.split()
        data.append((curTime,r[0][:-1],TranV4intoP((r[1])),TranV4intoP((r[2])),TranStateintoSTR(int(r[3],16))))
    f.close()
    TotalData["rawipv4"]=data
    f=open("/proc/net/raw6","r")
    data=[]
    for item in f.readlines()[1:]:
        r=item.split()
        data.append((curTime,r[0][:-1],TranV6intoP(r[1]),TranV6intoP(r[2]),TranStateintoSTR(int(r[3],16))))
    TotalData["rawipv6"]=data
    f.close()
    pass

def GetDevInfo():
    f=open("/proc/net/dev","r")
    data=[]
    for item in f.readlines()[2:]:
        r=item.split()
        data.append((curTime,r[0][:-1]))
    TotalData["dev"]=data
    f.close()
    pass

def ListAll():
    global curTime
    curTime=time.time()
    GetDevInfo()
    GetTcpInfo()
    GetUdpInfo()
    GetRawInfo()
    GetIcmpInfo()
    return json.dumps(TotalData)


if __name__ == "__main__":

    print(ListAll())
