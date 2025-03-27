# Authentication Methods: Beyond Single Sign-On

## 1. Multi-Factor Authentication (MFA)

### Overview

Multi-Factor Authentication adds additional layers of security by requiring two or more verification methods.

### Types of Factors

- Something you know (password)
- Something you have (mobile device)
- Something you are (biometrics)
- Somewhere you are (location)

### Implementation Example (Golang)

```go
package authentication

import (
    "crypto/rand"
    "encoding/base32"
    "time"
    "github.com/pquerna/otp/totp"
)

type MultiFactorAuth struct {
    Secret string
    User   string
}

func GenerateTOTPSecret(username string) (string, error) {
    // Generate a random secret key
    randomBytes := make([]byte, 20)
    _, err := rand.Read(randomBytes)
    if err != nil {
        return "", err
    }

    secret := base32.StdEncoding.EncodeToString(randomBytes)
    return secret, nil
}

func ValidateTOTPCode(secret, userCode string) bool {
    return totp.Validate(userCode, secret)
}
```

### Node.js MFA Implementation

```javascript
const speakeasy = require("speakeasy");
const QRCode = require("qrcode");

class MultiFactorAuthentication {
  generateSecret(username) {
    const secret = speakeasy.generateSecret({
      name: `MyApp:${username}`,
    });

    return {
      secret: secret.base32,
      qrCodeUrl: secret.otpauth_url,
    };
  }

  verifyToken(secret, token) {
    return speakeasy.totp.verify({
      secret: secret,
      encoding: "base32",
      token: token,
    });
  }
}
```

## 2. Token-Based Authentication

### Key Characteristics

- Stateless authentication
- Typically uses JSON Web Tokens (JWT)
- Scalable across distributed systems

### Golang JWT Implementation

```go
package tokenauth

import (
    "time"
    "github.com/golang-jwt/jwt"
)

type CustomClaims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func GenerateToken(username string, secretKey []byte) (string, error) {
    claims := CustomClaims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
            Issuer:    "AuthenticationService",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}
```

## 3. Biometric Authentication

### Types

- Fingerprint scanning
- Facial recognition
- Voice recognition
- Retina/iris scanning

### Simplified Node.js Biometric Interface

```javascript
class BiometricAuth {
  async validateFingerprint(userId, fingerprintData) {
    // Placeholder for biometric validation logic
    // In real-world scenario, this would interface with
    // specialized biometric hardware/software
    const registeredFingerprint = await this.getStoredFingerprint(userId);
    return this.compareFingerprints(registeredFingerprint, fingerprintData);
  }

  async compareFingerprints(stored, input) {
    // Complex comparison algorithm
    // Returns match probability
    return Math.random() > 0.5;
  }
}
```

## 4. Passwordless Authentication

### Methods

- Email magic links
- SMS verification codes
- Authentication apps
- Cryptographic tokens

### Golang Passwordless Example

```go
package passwordless

import (
    "crypto/rand"
    "encoding/hex"
    "time"
)

type PasswordlessAuth struct {
    Token        string
    ExpiresAt    time.Time
    User         string
}

func GenerateMagicLink(email string) (string, error) {
    tokenBytes := make([]byte, 16)
    _, err := rand.Read(tokenBytes)
    if err != nil {
        return "", err
    }

    token := hex.EncodeToString(tokenBytes)
    magicLink := fmt.Sprintf("https://myapp.com/auth/verify?token=%s", token)

    return magicLink, nil
}
```

## 5. OAuth and Social Authentication

### Providers

- Google
- Facebook
- GitHub
- LinkedIn

### Node.js Social Authentication Example

```javascript
const passport = require("passport");
const GoogleStrategy = require("passport-google-oauth20").Strategy;

passport.use(
  new GoogleStrategy(
    {
      clientID: GOOGLE_CLIENT_ID,
      clientSecret: GOOGLE_CLIENT_SECRET,
      callbackURL: "/auth/google/callback",
    },
    function (accessToken, refreshToken, profile, cb) {
      // Find or create user in database
      User.findOrCreate({ googleId: profile.id }, function (err, user) {
        return cb(err, user);
      });
    }
  )
);
```

## Security Considerations

1. Encrypt sensitive data
2. Implement proper error handling
3. Use secure, randomized tokens
4. Implement token expiration
5. Log authentication attempts
6. Use HTTPS
7. Implement rate limiting

## Comparative Analysis

| Method       | Pros                | Cons                      |
| ------------ | ------------------- | ------------------------- |
| MFA          | High security       | Complex implementation    |
| Token-Based  | Stateless, scalable | Token management          |
| Biometric    | Difficult to forge  | Privacy concerns          |
| Passwordless | User-friendly       | Potential delivery issues |
| OAuth        | Convenient          | Dependency on third-party |

## Conclusion

Authentication is a critical security component. The best approach often combines multiple methods tailored to specific use cases and security requirements.
