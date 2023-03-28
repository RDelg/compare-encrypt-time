import sys

from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes
from cryptography.hazmat.primitives import padding
from cryptography.hazmat.backends import default_backend
from base64 import b64decode, b64encode

backend = default_backend()
padder = padding.PKCS7(128).padder()
unpadder = padding.PKCS7(128).unpadder()

# data = b"demo"
# data = padder.update(data) + padder.finalize()
key = b64decode("HJkPmTz+uY7wd0p1+w//DABgbvPq9/230RwEG2sJ9mo=")
iv = b64decode("AAAAAAAAAAAAAAAAAAAAAA==")


for line in sys.stdin:
    data = line.strip().encode()
    padder = padding.PKCS7(128).padder()
    data = padder.update(data) + padder.finalize()
    cipher = Cipher(algorithms.AES(key), modes.CBC(iv), backend=backend)
    encryptor = cipher.encryptor()
    ct = encryptor.update(data) + encryptor.finalize()
    ct_out = b64encode(ct).decode()
    print(ct_out)
