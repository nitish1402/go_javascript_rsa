<!DOCTYPE html>
<html>
<body>
<script>
    async function decryptRSA(ciphertextBase64, privateKeyPem) {
        // Convert PEM encoded private key to ArrayBuffer
        privateKeyPem = privateKeyPem.replaceAll(/^\-+[^\-]+\-+$/gm, "").replace(/\n/gm, "");
        const binaryDerString = window.atob(privateKeyPem);
        const binaryDer = str2ab(binaryDerString);

        // Import the private key
        const privateKey = await window.crypto.subtle.importKey(
            "pkcs8",
            binaryDer,
            {
                name: "RSA-OAEP",
                hash: {name: "SHA-256"},
            },
            true,
            ["decrypt"]
        );

        // Decode the ciphertext from base64
        const ciphertext = Uint8Array.from(atob(ciphertextBase64), c => c.charCodeAt(0));

        // Decrypt the message
        const decrypted = await window.crypto.subtle.decrypt(
            {
                name: "RSA-OAEP"
            },
            privateKey,
            ciphertext
        );

        // Convert decrypted ArrayBuffer to string
        const decoder = new TextDecoder();
        console.log("Decrypted message:", decoder.decode(decrypted));
    }

    // Helper function to convert base64-encoded string to ArrayBuffer
    function str2ab(str) {
        const buf = new ArrayBuffer(str.length);
        const bufView = new Uint8Array(buf);
        for (let i = 0; i < str.length; i++) {
            bufView[i] = str.charCodeAt(i);
        }
        return buf;
    }

    // Example usage
    const ciphertext = "IXLQgqoto7Qw2G0YJAC7FwMw+LhconMKNsME/rb94IAT2XAL6YHCgrVQJqBhXWKiGOipP5vgpMwYo/s8rw0plJpL8WXtwPYQDU1+9T9w9PBWR0SRbZ/KLgFHZVk1RsgjS1mOKowH76Eba7v4M+7IpeMP4BOhB8DL7rB6IFuRSr/rXEFROyuGyvI1SQWB2J+YEOTlptAZQnLW0yaJBUDGu906LihGntTWdddgb7gkLYbWgzRHWY4SQzamOhi7rN8+qAaxA/qoSfCHAj26G3vebGk2nfDLPpsZeyMGJxi/UOSYtfvWMcKgFRuqRr4T9/BftyP3t1QOask5wupki0UPLQ=="; // Replace this with Golang output
    const privateKeyPem = `-----BEGIN PRIVATE KEY-----
    MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDCRvPjdIkfNOaZ
    E2gBq68S+Cuw8CblgJaxfelB4CouKO07VbHYJ9ZMi2co95rZ/ilnRPgm0NibTRFb
    6G0Qr45JTuqCdE6Mbz0L1fM0iqMRiq/+cnf63EOKif5+s3UfE1YZNRcZiI86K+v/
    BbviOM4mVdHg+KK1yK5RjpztwhJWZidZzZY20nfDbF1LeN/aFIUF4h0Pz+Rjt0cO
    Y6T9cUBJKCNolEt6zs6D8sABRh6h3y5tCydZAZ50lnFRMPjZR9IuH7jElggub68L
    dA8QTXc/odHpinS6qdDjYb5L7+yo6DQ8kW6tf6h9pWHU62bi9Vm1CyajOmuIEh5o
    dNQ2zw+3AgMBAAECggEABXyxiLwS32HEHZCxwFJUCIof2ch+oL7IAs1WmDi3mEQp
    pyJdeohtgg3x5PNRWVTXYIZPJ/GAHKrJkbn30p/SuflToEmeqlK9+6aYTuSXhHhR
    TjN3dgtgkPoiyPtSlIUcSmv4cg0sVm1FJhfIXbRTBjwoSF1dYxr+6WjIv+JaWDxP
    xt/cpHcLX2kEdnxWMgboqDZH7mkKCIVE/r2Dm8kngm9Qw+S9oVmgBDjvz8CkfkCV
    qlZfmsmKjbogxh8VAPVrNI762zN1Z1Q/fkJcfNxgnyYM4PNOOKBgRWzCp8M6o+8h
    /gqyuNcQz2/3Hk+anbdlHYzujDaUXcpO+QJUMg8S2QKBgQDntA266RS0eI1h+Y/u
    mxYSq5zhSEsd0xEUUS9f/tgv35aelDfG2/uAQUsOU4wF0EWDRto5N9idnE8NC5fy
    dFh/KpZFUZVKHWHYbkfvjwtHvMEDtb6nObFoC0zxvmnXv9GpoFY0Vc9CgBm8G5iP
    spqby4qjoMFKWX2vihz7eAzaGwKBgQDWpjbdVUDCQu8K4PGPsRR7DI1ofn+Y2PCs
    KxXLYlj9G1AS1qSAdW/z+B2pj8ph/vP+APtqKmgFyWLYRiu/cFsKRO3NUlHX0gkp
    HG6LNnR2HJQudigHDCjUeMkiyC40AZljCAM/YqT3sV8JZxJDTEpOldJfuODG9PPB
    zojqiF46lQKBgQCEcgHfM9joCHkY5iUGSZRme76jcEWv+LSsnnOsNeqyAucAIs13
    WMv81lXnDI7fy9vQXLHlPy0Newoc9OGYcDUeC+P2H9pskTsbEn8ASw4xpY3XZw5i
    XLIyPTNwhF3QPA/HHKXqKJre8obDZirhCUEjiUfonL2gmGMl8pb6j/cYvQKBgQCc
    v+bi/SX8dwq/xANDrspJDaKag35ErXAcGp6g1uIre+2exDZ/RMOLw05ODB58L8YY
    YLk5D1zFlQpk6+F4vOEO/8U1In93/v8Hkaa8bPjhY/9maozSkLOkbcxcZRkwi1Zr
    NmfAuxXDeDjx0d5JXQYKm2h01Lr5L9puVpgvxYQifQKBgCmTOsl0i/A+KF2bDCtE
    NUh5kQh21KBSktos8fKS9bOjva/BFQuDr/2CeDw6txo91p9n7Cevka/TQDo+0gpC
    ivVbAodzc7OT4bwB+Ti7wAx/kc7hyLvds3dKB3IvPRxr/KNxER6uUbKlzqtMXnKZ
    1qCijkA2OH+K0ut2Sr54avqL
    -----END PRIVATE KEY-----`; // Replace with the PEM encoded private key

    decryptRSA(ciphertext, privateKeyPem);



    <!DOCTYPE html>
<html>
<body>
<script>
    async function encryptRSA(plaintext, publicKeyPem) {
        // Convert PEM encoded public key to ArrayBuffer

        publicKeyPem = publicKeyPem.replaceAll(/^\-+[^\-]+\-+$/gm, "").replace(/\n/gm, "");
        const binaryDerString = window.atob(publicKeyPem);
        const binaryDer = str2ab(binaryDerString);

        // Import the public key
        const publicKey = await window.crypto.subtle.importKey(
            "spki",  // "spki" for public key format
            binaryDer,
            {
                name: "RSA-OAEP",
                hash: {name: "SHA-256"},
            },
            true,
            ["encrypt"]
        );

        // Encode the plaintext into an ArrayBuffer
        const encoder = new TextEncoder();
        const encoded = encoder.encode(plaintext);

        // Encrypt the data
        const ciphertext = await window.crypto.subtle.encrypt(
            {
                name: "RSA-OAEP"
            },
            publicKey,
            encoded
        );

        // Convert ciphertext to base64 for easier transmission
        const base64Ciphertext = btoa(String.fromCharCode(...new Uint8Array(ciphertext)));
        console.log("Encrypted message (Base64):", base64Ciphertext);
        return base64Ciphertext;
    }

    // Helper function to convert base64-encoded string to ArrayBuffer
    function str2ab(str) {
        const buf = new ArrayBuffer(str.length);
        const bufView = new Uint8Array(buf);
        for (let i = 0; i < str.length; i++) {
            bufView[i] = str.charCodeAt(i);
        }
        return buf;
    }

    // Example usage
    const publicKeyPem = `-----BEGIN PUBLIC KEY-----
    MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwkbz43SJHzTmmRNoAauv
    EvgrsPAm5YCWsX3pQeAqLijtO1Wx2CfWTItnKPea2f4pZ0T4JtDYm00RW+htEK+O
    SU7qgnROjG89C9XzNIqjEYqv/nJ3+txDion+frN1HxNWGTUXGYiPOivr/wW74jjO
    JlXR4PiitciuUY6c7cISVmYnWc2WNtJ3w2xdS3jf2hSFBeIdD8/kY7dHDmOk/XFA
    SSgjaJRLes7Og/LAAUYeod8ubQsnWQGedJZxUTD42UfSLh+4xJYILm+vC3QPEE13
    P6HR6Yp0uqnQ42G+S+/sqOg0PJFurX+ofaVh1Otm4vVZtQsmozpriBIeaHTUNs8P
    twIDAQAB
    -----END PUBLIC KEY-----`;  // Replace this with actual PEM encoded public key
    const plaintext = "Hello, RSA Encryption from JavaScript!";

    encryptRSA(plaintext, publicKeyPem);
</script>
</body>
</html>

</script>
</body>
</html>
