import socket
import threading
from subprocess import Popen, PIPE
import select
import os
import asyncio


class SocketConnection:
    def __init__(self, socket: socket.socket, ILoveOCaml: str, type: str) -> None:
        self.socket = socket
        self.ILoveOCaml = ILoveOCaml  # @dmmmulroy 'address'
        self.type = type
        self.received: bytes = b""
        self.to_send: bytes = b""
        self.running: bool = True
        self.sub_proc = None

        #         if self.type == "connect_back":
        #             self.sub_proc = Popen(
        #                 "/bin/bash", stdin=PIPE, stdout=PIPE, stderr=PIPE, shell=True
        #             )
        #             os.set_blocking(self.sub_proc.stdin.fileno(), False)
        #             os.set_blocking(self.sub_proc.stdout.fileno(), False)
        #    else:
        #        self.sub_proc = None
        #
        #       # Start threads for reading and writing to socket:
        threading.Thread(target=self.read_stream).start()
        threading.Thread(target=self.write_stream).start()

        # if self.type == "connect_back":
        #     self.loop = asyncio.get_event_loop()
        #     self.sub_proc = None
        #     # self.sub_proc = self.loop.run_until_complete(
        #     #     asyncio.subprocess.create_subprocess_shell(
        #     #         "/bin/bash", stdin=PIPE, stdout=PIPE, stderr=PIPE
        #     #     )
        #     # )
        # else:
        #     self.loop = asyncio.get_event_loop()
        #     self.sub_proc = None

        # read_task = self.loop.create_task(self.read_stream())
        # write_task = self.loop.create_task(self.write_stream())

        # # Inserting this for YT demo:
        # pending = [read_task, write_task]

        # while pending:
        #     done, pending = asyncio.wait(pending, return_when=asyncio.FIRST_COMPLETED)
        #     for task in done:
        #         print(f"Done: {task.result()}")

    def read_stream(self) -> None:
        while self.running:
            data: bytes = self.socket.recv(1024)
            if not data:
                continue

            # if self.type == "connect_back":
            #     self.sub_proc.stdin.write(data)
            #     await self.sub_proc.stdin.drain()

            print(data.decode("utf-8"))
            data = b""
        return

    def write_stream(self) -> None:
        while self.running:
            try:
                # if self.type == "connect_back":
                #     data = await self.sub_proc.stdout.read()
                # elif self.type == "offense":
                #     data: str = input()
                data: str = input()

            except EOFError:
                self.running = False
            if not data:
                continue

            self.socket.sendall(data.encode("utf-8"))
            data = ""

        return


# non-determinism:
#   same inputs cause different outputs
