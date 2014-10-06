zipServer

Serve up a whole website from a zip file, without decompressing to intermediate storage.

Build:
go build zipServer.go


Usage:
Let it be that a zip file "voynich_manuscript.zip" exists, whose contents, including a
file "index.html", are inside a folder by the name of "vman"

	on server (ipaddress == 192.168.1.1):
	./zipServer ./voynich_manuscript.zip

	on client:
	http://192.168.1.1/vman/index.html


Limitations:
This is a horrible hack by a someone who doesn't know much Go.
Thus there are obviously some limitations, most of which are probably unknown.
Of the known:
-There seems to be a limitation to zipfs that prevents mounting the "root" directory
of a zip file.  Only subfolders can be mounted.


License Stuff:

In the "vfs" folder you'll find the code that's been re-purposed from the godoc
tool, as well as the AUTHORS, LICENSE, etc. files for that code.

Also, you'll find here a .zip file ("voynich_manuscript.zip"), a thumbnail copy of the
Voynich Manuscript gotten from Archive.org, just for testing purposes. That is a
Public Domain work.

As for my paltry code and modifications:

zipServer
(C) 2014, Zilog8 <zeuscoding@gmail.com>

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice,
   this list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.
3. The name of the author may not be used to endorse or promote products
   derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR IMPLIED
WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO
EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
