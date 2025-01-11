# 🔐 Fundamentos de Cifrado

## 1️⃣ Cifrado Simétrico (AES-GCM)

* __Clave Única__: Se usa una misma clave para cifrar y descifrar.
* __Rápido y Eficiente__: Ideal para cifrar grandes volúmenes de datos. [Ver benchmark](https://medium.com/@gerritjvv/aes-golang-encryption-performance-benchmarks-updated-bcfa3555165b)
* __Seguridad__: Depende de mantener la clave secreta.

### 🔎 Detalles del AES-GCM:

* __AES (Advanced Encryption Standard)__: Algoritmo de cifrado por bloques.
* __GCM (Galois/Counter Mode)__: Proporciona autenticación e integridad de los datos.
* __Nonce (Número Único)__: Garantiza que los datos cifrados sean únicos aunque el mensaje sea el mismo.

### 📦 Uso:

* Cifrado de archivos, bases de datos, almacenamiento local.

---

## 2️⃣ Cifrado Asimétrico (RSA)

* __Clave Pública y Privada__:
  * __Clave Pública__: Para cifrar.
  * __Clave Privada__: Para descifrar.
* __Seguridad Basada en Matemáticas__: Difícil de romper por factorización de números grandes.
* __Más lento__, pero ideal para compartir datos seguros.

### 📦 Uso:

* __TLS/SSL__ (_HTTPS_), autenticación de identidad, intercambio de claves.


⚠️ __RSA no se usa para cifrar datos grandes__, sino para cifrar claves simétricas o datos pequeños.

---

## 🔑 Gestión de Claves

### 🔒 AES (Clave Simétrica)

* __Generación__: Se genera con suficiente aleatoriedad (`crypto/rand`)
* __Almacenamiento__:
  * __Archivo seguro__ (`0600`).
  * __Variable de entorno__.
  * __Gestores de secretos__ (_AWS KMS, Vault_).
* __Rotación__: Cambiar periódicamente la clave para mayor seguridad.

### 🔐 RSA (Clave Asimétrica)

* __Generación__: Par de claves (privada y pública).
* __Almacenamiento__:
  * __Clave privada__ → protegida (`0600`).
  * __Clave pública__ → se puede compartir libremente.
* __Exportación__: Se usa formato __PEM__ (_legible_) para compartir e importar claves.

---

## 📊 Comparación AES vs RSA

| Característica | AES (_Simétrico_)              | RSA (_Asimétrico_)         |
|----------------|--------------------------------|----------------------------|
| __Velocidad__  | ⚡ Muy rápido                   | 🐢 Más lento               |
| __Clave__      | Única clave secreta            | Clave pública y privada    |
| __Eficiencia__ | Mejor para grandes volúmenes   | Mejor para datos pequeños  |
| __Seguridad__  | Depende del manejo de la clave | Seguridad matemática       |
| __Uso común__  | Almacenamiento, bases de datos | HTTPS, firmas digitales    |

### 🔎 ¿Por qué usar ambos juntos?
* RSA cifra la clave AES, y AES cifra los _datos grandes_. Esto combina lo mejor de ambos mundos.

---

## 🛠️ Buenas Prácticas
1. __Protege las claves privadas__: Usa permisos (`0600`) y evita compartirlas.
2. __Usa AEAD (AES-GCM)__: Para integridad y autenticación de los datos.
3. __Rota claves periódicamente__: No uses la misma clave por siempre.
4. __Evita hardcodear claves__: Usa variables de entorno o gestores de secretos.
5. __No reutilices Nonces__: Siempre genera un nonce único para cada cifrado.

---

## 🏆 Resumen

* __AES (Simétrico)__: Rápido y eficiente, pero debes proteger la clave.
* __RSA (Asimétrico)__: Seguro para compartir información, pero más lento.
* __Mejor práctica__: Usar RSA para cifrar la clave AES y AES para los datos.
