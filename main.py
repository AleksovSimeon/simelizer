import ctypes
import json

library = ctypes.cdll.LoadLibrary("./library.so")
hello_world = library.helloWorld
hello_world()

hello = library.hello
hello.argtypes = [ctypes.c_char_p]
hello("simeon".encode("utf-8"))
