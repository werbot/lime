# Lime License Server - API Documentation

## Table of Contents

1. [Overview](#overview)
2. [Backend APIs](#backend-apis)
   - [Public API Endpoints](#public-api-endpoints)
   - [Admin API Endpoints](#admin-api-endpoints)
   - [Manager API Endpoints](#manager-api-endpoints)
3. [Go Packages](#go-packages)
   - [fsutil Package](#fsutil-package)
   - [jwtutil Package](#jwtutil-package)
   - [geo Package](#geo-package)
   - [security Package](#security-package)
   - [strutil Package](#strutil-package)
   - [webutil Package](#webutil-package)
   - [crypto Package](#crypto-package)
   - [logging Package](#logging-package)
   - [storage Package](#storage-package)
   - [archive Package](#archive-package)
   - [license Package](#license-package)
4. [Frontend Components](#frontend-components)
   - [Vue Components](#vue-components)
   - [Utility Functions](#utility-functions)
   - [API Client](#api-client)
5. [Application Structure](#application-structure)
6. [Configuration](#configuration)
7. [Examples](#examples)

## Overview

Lime is a lightweight license key server built with Go (backend) and Vue.js (frontend). It provides a complete solution for license management, customer management, and payment processing.

### Key Features
- License key generation and validation
- Customer management
- Payment processing (Stripe integration)
- Geographic location detection
- JWT-based authentication
- Admin and public interfaces

## Backend APIs

### Public API Endpoints

#### Health Check
```http
GET /ping
```
Returns a simple health check response.

**Response:**
```json
{
  "status": "ok"
}
```

#### License Download
```http
GET /api/license/download/:id
```
Downloads a license file by ID.

**Parameters:**
- `id` (path): License ID (format: `lic_` or `cst_` followed by 15 characters)

**Response:** License file download

#### License Verification
```http
POST /api/license/verify
```
Verifies a license key.

**Request Body:**
```json
{
  "license_key": "your_license_key_here"
}
```

**Response:**
```json
{
  "valid": true,
  "license": {
    "id": "lic_abc123def456ghi",
    "customer_id": "cst_xyz789uvw012abc",
    "pattern_id": "pat_123456789abcdef",
    "status": "active",
    "expires_at": "2024-12-31T23:59:59Z"
  }
}
```

### Admin API Endpoints

All admin endpoints require JWT authentication with "admin" credentials.

#### Authentication
```http
POST /_/api/sign/in
```
Sign in to admin panel.

**Request Body:**
```json
{
  "email": "admin@example.com",
  "password": "password"
}
```

**Response:**
```json
{
  "token": "jwt_token_here",
  "user": {
    "id": "user_id",
    "email": "admin@example.com"
  }
}
```

```http
POST /_/api/sign/out
```
Sign out from admin panel.

#### License Management
```http
GET /_/api/license/
```
Get all licenses.

**Query Parameters:**
- `page`: Page number
- `limit`: Items per page
- `search`: Search term

```http
POST /_/api/license/
```
Create a new license.

**Request Body:**
```json
{
  "customer_id": "cst_xyz789uvw012abc",
  "pattern_id": "pat_123456789abcdef",
  "expires_at": "2024-12-31T23:59:59Z",
  "metadata": {
    "domain": "example.com"
  }
}
```

```http
GET /_/api/license/:id
```
Get license by ID.

**Parameters:**
- `id` (path): License ID (format: `lic_` or `cst_` followed by 15 characters)

```http
PATCH /_/api/license/:id
```
Update license.

**Parameters:**
- `id` (path): License ID (15 characters)

**Request Body:**
```json
{
  "status": "suspended",
  "metadata": {
    "domain": "newdomain.com"
  }
}
```

#### Pattern Management
```http
GET /_/api/pattern/
```
Get all patterns.

```http
POST /_/api/pattern/
```
Create a new pattern.

**Request Body:**
```json
{
  "name": "Premium License",
  "description": "Premium features included",
  "price": 9999,
  "currency": "USD",
  "duration": 365,
  "duration_unit": "day"
}
```

```http
POST /_/api/pattern/:id
```
Clone a pattern.

**Parameters:**
- `id` (path): Pattern ID (15 characters)

```http
GET /_/api/pattern/:id
```
Get pattern by ID.

```http
PATCH /_/api/pattern/:id
```
Update pattern.

```http
DELETE /_/api/pattern/:id
```
Delete pattern.

#### Customer Management
```http
GET /_/api/customer/
```
Get all customers.

```http
POST /_/api/customer/
```
Create a new customer.

**Request Body:**
```json
{
  "email": "customer@example.com",
  "name": "John Doe",
  "company": "Example Corp"
}
```

```http
GET /_/api/customer/:id
```
Get customer by ID.

```http
PATCH /_/api/customer/:id
```
Update customer.

```http
DELETE /_/api/customer/:id
```
Delete customer.

#### Payment Management
```http
GET /_/api/payment/
```
Get all payments.

**Query Parameters:**
- `filter`: Filter type (`no_lic` or `all`)

```http
POST /_/api/payment/
```
Create a new payment.

```http
GET /_/api/payment/:id
```
Get payment by ID.

```http
PATCH /_/api/payment/:id
```
Update payment.

#### Audit Logs
```http
GET /_/api/audit/
```
Get audit logs.

```http
GET /_/api/audit/:id
```
Get audit log by ID.

#### Lists
```http
GET /_/api/list/patterns/:pattern?
```
Get patterns list.

```http
GET /_/api/list/customers/:customer?
```
Get customers list.

```http
GET /_/api/list/countries/:country
```
Get countries list.

#### Settings
```http
GET /_/api/setting/:group
```
Get settings by group.

```http
PATCH /_/api/setting/:group
```
Update settings group.

### Manager API Endpoints

Manager endpoints are similar to admin endpoints but with different authentication requirements.

## Go Packages

### fsutil Package

File system utility functions for common file and directory operations.

#### Constants
```go
const (
    FsCWAFlags = os.O_CREATE | os.O_WRONLY | os.O_APPEND // create, append write-only
    FsCWTFlags = os.O_CREATE | os.O_WRONLY | os.O_TRUNC  // create, override write-only
    FsCWFlags  = os.O_CREATE | os.O_WRONLY               // create, write-only
    FsRFlags   = os.O_RDONLY                             // read-only
)
```

#### Functions

##### IsFile
```go
func IsFile(path string) bool
```
Reports whether the named file exists and is not a directory.

**Parameters:**
- `path`: File path to check

**Returns:** `true` if path exists and is a file, `false` otherwise

**Example:**
```go
if fsutil.IsFile("/path/to/file.txt") {
    fmt.Println("File exists")
}
```

##### MustReadFile
```go
func MustReadFile(filePath string) []byte
```
Reads file contents and panics on error.

**Parameters:**
- `filePath`: Path to the file to read

**Returns:** File contents as byte slice

**Example:**
```go
content := fsutil.MustReadFile("/path/to/file.txt")
```

##### OpenFile
```go
func OpenFile(filepath string, flag int, perm os.FileMode) (*os.File, error)
```
Opens a file like `os.OpenFile` but automatically creates the directory if it doesn't exist.

**Parameters:**
- `filepath`: Path to the file
- `flag`: File opening flags
- `perm`: File permissions

**Returns:** File handle and error

**Example:**
```go
file, err := fsutil.OpenFile("/path/to/file.txt", fsutil.FsCWTFlags, 0644)
if err != nil {
    log.Fatal(err)
}
defer file.Close()
```

##### WriteOSFile
```go
func WriteOSFile(f *os.File, data any) (n int, err error)
```
Writes data to a file and closes it. Supports string, []byte, and io.Reader types.

**Parameters:**
- `f`: File handle
- `data`: Data to write (string, []byte, or io.Reader)

**Returns:** Number of bytes written and error

**Example:**
```go
file, _ := os.Create("test.txt")
n, err := fsutil.WriteOSFile(file, "Hello, World!")
```

##### ExtName
```go
func ExtName(fpath string) string
```
Extracts the file extension without the dot.

**Parameters:**
- `fpath`: File path

**Returns:** File extension

**Example:**
```go
ext := fsutil.ExtName("file.txt") // Returns "txt"
```

##### CopyFile
```go
func CopyFile(srcPath, dstPath string) error
```
Copies a file from source to destination.

**Parameters:**
- `srcPath`: Source file path
- `dstPath`: Destination file path

**Returns:** Error if any

**Example:**
```go
err := fsutil.CopyFile("source.txt", "destination.txt")
```

##### IsDir
```go
func IsDir(path string) bool
```
Reports whether the named directory exists.

**Parameters:**
- `path`: Directory path to check

**Returns:** `true` if path exists and is a directory

**Example:**
```go
if fsutil.IsDir("/path/to/directory") {
    fmt.Println("Directory exists")
}
```

##### Workdir
```go
func Workdir() string
```
Gets the current working directory.

**Returns:** Current working directory path

**Example:**
```go
wd := fsutil.Workdir()
fmt.Println("Working directory:", wd)
```

##### MkDirs
```go
func MkDirs(perm os.FileMode, dirPaths ...string) error
```
Creates multiple directories at once.

**Parameters:**
- `perm`: Directory permissions
- `dirPaths`: Variable number of directory paths

**Returns:** Error if any

**Example:**
```go
err := fsutil.MkDirs(0755, "/path/dir1", "/path/dir2", "/path/dir3")
```

##### MkSubDirs
```go
func MkSubDirs(perm os.FileMode, parentDir string, subDirs ...string) error
```
Creates multiple subdirectories under a parent directory.

**Parameters:**
- `perm`: Directory permissions
- `parentDir`: Parent directory path
- `subDirs`: Variable number of subdirectory names

**Returns:** Error if any

**Example:**
```go
err := fsutil.MkSubDirs(0755, "/parent", "sub1", "sub2", "sub3")
```

##### RemoveDir
```go
func RemoveDir(dir string) error
```
Removes a directory and all its contents.

**Parameters:**
- `dir`: Directory path to remove

**Returns:** Error if any

**Example:**
```go
err := fsutil.RemoveDir("/path/to/remove")
```

##### Download
```go
func Download(filepath string, url string) (err error)
```
Downloads a file from a URL and saves it to the specified filepath.

**Parameters:**
- `filepath`: Local file path to save the download
- `url`: URL to download from

**Returns:** Error if any

**Example:**
```go
err := fsutil.Download("local_file.zip", "https://example.com/file.zip")
```

### jwtutil Package

JWT utility functions for token generation, parsing, and validation.

#### Types

##### TokenMetadata
```go
type TokenMetadata struct {
    ID        string
    IssuedAt  float64
    ExpiresAt float64
}
```

#### Variables
```go
var (
    ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
    ErrVerificationError       = errors.New("verification error")
    ErrInvalidToken            = errors.New("the token is invalid")
)
```

#### Functions

##### LoadKeys
```go
func LoadKeys(pubKeyPath, privKeyPath string) error
```
Loads RSA public and private keys from PEM files.

**Parameters:**
- `pubKeyPath`: Path to public key file
- `privKeyPath`: Path to private key file

**Returns:** Error if any

**Example:**
```go
err := jwtutil.LoadKeys("public.pem", "private.pem")
```

##### PublicKey
```go
func PublicKey() *rsa.PublicKey
```
Returns the loaded public key.

**Returns:** RSA public key

##### PrivateKey
```go
func PrivateKey() *rsa.PrivateKey
```
Returns the loaded private key.

**Returns:** RSA private key

##### NewToken
```go
func NewToken(id, expires string, credentials []string) (string, error)
```
Generates a new JWT access token.

**Parameters:**
- `id`: User ID
- `expires`: Token expiration duration (e.g., "24h", "7d")
- `credentials`: Array of user credentials/permissions

**Returns:** JWT token string and error

**Example:**
```go
token, err := jwtutil.NewToken("user123", "24h", []string{"admin", "read"})
```

##### ExtractMetadata
```go
func ExtractMetadata(tokenKey string) (*TokenMetadata, error)
```
Extracts metadata from a JWT token.

**Parameters:**
- `tokenKey`: JWT token string

**Returns:** Token metadata and error

**Example:**
```go
metadata, err := jwtutil.ExtractMetadata(tokenString)
if err == nil {
    fmt.Printf("User ID: %s\n", metadata.ID)
}
```

##### ExtractMetadataFiber
```go
func ExtractMetadataFiber(c *fiber.Ctx) (*TokenMetadata, error)
```
Extracts metadata from a JWT token within a Fiber context.

**Parameters:**
- `c`: Fiber context

**Returns:** Token metadata and error

**Example:**
```go
func handler(c *fiber.Ctx) error {
    metadata, err := jwtutil.ExtractMetadataFiber(c)
    if err != nil {
        return err
    }
    // Use metadata
    return nil
}
```

### geo Package

Geographic location utilities for IP-based country detection.

#### Types

##### Storage
```go
type Storage string

const (
    Maxmind Storage = "maxmind"
    Ipinfo  Storage = "ipinfo"
    GeoOpen Storage = "geoopen"
)
```

##### Database
```go
type Database struct {
    DBPath  string         `toml:"db-path"`
    Storage Storage        `toml:"storage"`
    GeoOpen geoopen.Config `toml:"geoopen"`
    Maxmind maxmind.Config `toml:"maxmind,commented"`
    Ipinfo  ipinfo.Config  `toml:"ipinfo,commented"`
}
```

#### Methods

##### Download
```go
func (db *Database) Download() error
```
Downloads the geographic database based on the configured storage type.

**Returns:** Error if any

**Example:**
```go
db := &geo.Database{
    DBPath:  "/path/to/db",
    Storage: geo.GeoOpen,
}
err := db.Download()
```

##### Check
```go
func (db *Database) Check() bool
```
Checks if the geographic database exists and is valid.

**Returns:** `true` if database is available

**Example:**
```go
if db.Check() {
    fmt.Println("Database is ready")
}
```

### security Package

Security utilities for password hashing and token generation.

#### Constants
```go
const (
    DefaultIdLength   = 15
    DefaultIdAlphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)
```

#### Functions

##### NormalizePassword
```go
func NormalizePassword(p string) []byte
```
Normalizes a password string to byte slice.

**Parameters:**
- `p`: Password string

**Returns:** Password as byte slice

**Example:**
```go
pwdBytes := security.NormalizePassword("mypassword")
```

##### GeneratePassword
```go
func GeneratePassword(p string) string
```
Generates a bcrypt hash from a password.

**Parameters:**
- `p`: Plain text password

**Returns:** Bcrypt hash string

**Example:**
```go
hash := security.GeneratePassword("mypassword")
```

##### ComparePasswords
```go
func ComparePasswords(hashedPwd, inputPwd string) bool
```
Compares a plain text password with a bcrypt hash.

**Parameters:**
- `hashedPwd`: Bcrypt hash
- `inputPwd`: Plain text password to verify

**Returns:** `true` if passwords match

**Example:**
```go
if security.ComparePasswords(hashedPassword, inputPassword) {
    fmt.Println("Password is correct")
}
```

##### NewToken
```go
func NewToken(text string) (string, error)
```
Generates a new token from text using bcrypt and MD5.

**Parameters:**
- `text`: Text to generate token from

**Returns:** Token string and error

**Example:**
```go
token, err := security.NewToken("some text")
```

##### NanoID
```go
func NanoID() string
```
Generates a random ID using the default alphabet and length.

**Returns:** Random ID string

**Example:**
```go
id := security.NanoID() // e.g., "a1b2c3d4e5f6g7h"
```

### strutil Package

String utility functions for common string operations.

#### Functions

##### ToSlice
```go
func ToSlice(s string, sep ...string) []string
```
Splits a string into a slice using the specified separator.

**Parameters:**
- `s`: String to split
- `sep`: Optional separator (defaults to ",")

**Returns:** String slice

**Example:**
```go
// Using default separator (comma)
parts := strutil.ToSlice("a,b,c") // ["a", "b", "c"]

// Using custom separator
parts := strutil.ToSlice("a|b|c", "|") // ["a", "b", "c"]
```

##### ToAny
```go
func ToAny(key ...string) []any
```
Converts string arguments to interface{} slice.

**Parameters:**
- `key`: Variable number of strings

**Returns:** Interface slice

**Example:**
```go
result := strutil.ToAny("key1", "key2", "key3")
// result is []any{"key1", "key2", "key3"}
```

##### ToInt32
```go
func ToInt32(s string) int32
```
Converts a string to int32, returns 256 on error.

**Parameters:**
- `s`: String to convert

**Returns:** Int32 value

**Example:**
```go
num := strutil.ToInt32("123") // 123
num := strutil.ToInt32("abc") // 256 (default on error)
```

## Frontend Components

### Vue Components

#### Alert Component
```vue
<Alert />
```
Displays notification alerts using the Notiwind library.

**Features:**
- Supports multiple notification types (success, error, warning, info)
- Auto-dismiss after 4 seconds
- Responsive design
- Event-driven notifications

**Usage:**
```javascript
// Trigger notifications
window.dispatchEvent(new CustomEvent('connextSuccess', { 
  detail: 'Operation completed successfully!' 
}));

window.dispatchEvent(new CustomEvent('connextError', { 
  detail: 'An error occurred!' 
}));
```

#### Form Components

##### FormInput
```vue
<FormInput
  v-model="value"
  name="Field Name"
  id="field-id"
  type="text"
  :required="true"
  placeholder="Enter value"
  rules="required|email"
  :error="errorMessage"
/>
```

**Props:**
- `modelValue`: Input value (v-model)
- `name`: Field label
- `id`: Field ID
- `type`: Input type (default: "text")
- `class`: CSS classes
- `disabled`: Disable input
- `required`: Mark as required
- `placeholder`: Placeholder text
- `autocomplete`: Autocomplete attribute
- `rules`: Validation rules
- `error`: Error message

##### FormSelect
```vue
<FormSelect
  v-model="selectedValue"
  name="Select Option"
  :options="options"
  :required="true"
/>
```

##### FormToggle
```vue
<FormToggle
  v-model="enabled"
  name="Enable Feature"
/>
```

#### Navigation Component
```vue
<Navigation />
```
Provides navigation menu for the application.

#### Header Component
```vue
<Header />
```
Application header with user information and navigation.

#### Pagination Component
```vue
<Pagination
  :current-page="currentPage"
  :total-pages="totalPages"
  :total-items="totalItems"
  @page-change="handlePageChange"
/>
```

#### Badge Component
```vue
<Badge :type="'success'" :text="'Active'" />
```

#### Drawer Component
```vue
<Drawer v-model="isOpen" title="Drawer Title">
  <p>Drawer content</p>
</Drawer>
```

#### Tabs Component
```vue
<Tabs :tabs="tabs" v-model="activeTab">
  <template #tab1>Content for tab 1</template>
  <template #tab2>Content for tab 2</template>
</Tabs>
```

#### Skeleton Component
```vue
<Skeleton :lines="3" />
```

#### SvgIcon Component
```vue
<SvgIcon name="icon-name" class="w-6 h-6" />
```

### Utility Functions

#### Cookie Management
```javascript
// Get cookie value
const token = getCookie('auth_token');

// Set cookie
setCookie('auth_token', 'value', 7); // 7 days

// Delete cookie
delCookie('auth_token');
```

#### Price Formatting
```javascript
// Format price from cents to dollars
const formatted = priceFormat('9999'); // "99.99"

// Format cost
const cost = costFormat('9999'); // "99.99"

// Convert to Stripe format (cents)
const stripeAmount = costStripe(99.99); // 9999
```

#### Date Formatting
```javascript
// Format timestamp
const formatted = formatDate(1640995200); // "1 Jan 2022, 12:00:00 AM"
```

#### String Utilities
```javascript
// Generate random string
const random = randomString(10); // "aB3kL9mN2p"

// Capitalize first letter
const capitalized = firstLetter("hello"); // "Hello"
```

#### Array Utilities
```javascript
// Convert array to object
const obj = arrayToObject(['a', 'b', 'c'], item => item.toUpperCase());
// Result: { 1: 'A', 2: 'B', 3: 'C' }

// Reduce items to object
const reduced = reduceToObject(
  [{ key: 'name', value: 'John' }, { key: 'age', value: '30' }],
  value => value
);
// Result: { name: 'John', age: '30' }
```

### API Client

The API client provides HTTP request utilities for communicating with the backend.

#### Functions

##### apiGet
```javascript
const response = await apiGet('/api/endpoint', { param1: 'value1' });
```

##### apiPost
```javascript
const response = await apiPost('/api/endpoint', { param1: 'value1' }, { data: 'body' });
```

##### apiUpdate
```javascript
const response = await apiUpdate('/api/endpoint', { param1: 'value1' }, { data: 'body' });
```

##### apiDelete
```javascript
const response = await apiDelete('/api/endpoint', { param1: 'value1' });
```

**Features:**
- Automatic progress indicators
- Query parameter handling
- JSON body serialization
- Automatic 401 redirect handling
- Credentials inclusion

## Application Structure

### Backend Structure
```
├── cmd/                    # Command line applications
├── internal/               # Internal application code
│   ├── app.go             # Main application setup
│   ├── config/            # Configuration management
│   ├── errors/            # Error definitions
│   ├── handlers/          # HTTP request handlers
│   ├── middleware/        # HTTP middleware
│   ├── models/            # Data models
│   ├── queries/           # Database queries
│   └── routes/            # Route definitions
├── pkg/                   # Public packages
│   ├── fsutil/            # File system utilities
│   ├── jwtutil/           # JWT utilities
│   ├── geo/               # Geographic utilities
│   ├── security/          # Security utilities
│   ├── strutil/           # String utilities
│   └── webutil/           # Web utilities
├── web/                   # Frontend application
│   ├── src/               # Source code
│   │   ├── components/    # Vue components
│   │   ├── pages/         # Page components
│   │   ├── utils/         # Utility functions
│   │   └── store/         # State management
│   └── public/            # Static assets
└── migrations/            # Database migrations
```

### Frontend Structure
```
web/src/
├── components/            # Reusable Vue components
├── pages/                 # Page components
├── layouts/               # Layout components
├── store/                 # Pinia state management
├── utils/                 # Utility functions
│   ├── api/               # API client
│   └── menu/              # Menu configuration
└── assets/                # Static assets
```

## Configuration

### Environment Variables
- `HTTP_ADDR`: Server address (default: ":8088")
- `DEV_MODE`: Development mode flag
- `DATABASE_URL`: Database connection string
- `JWT_SECRET`: JWT secret key
- `GEO_DB_PATH`: Geographic database path

### Configuration File
The application uses TOML configuration files for settings:

```toml
[server]
http_addr = ":8088"
dev_mode = false

[database]
url = "sqlite:///lime.db"

[keys]
key_dir = "./keys"

[keys.jwt]
public_key = "jwt.pub"
private_key = "jwt.key"

[keys.license]
public_key = "license.pub"
private_key = "license.key"

[geo_database]
db_path = "./geo"
storage = "geoopen"
```

## Examples

### Complete License Management Example

#### Backend (Go)
```go
package main

import (
    "log"
    "github.com/werbot/lime/internal/app"
)

func main() {
    if err := app.NewApp(); err != nil {
        log.Fatal(err)
    }
}
```

#### Frontend (Vue.js)
```vue
<template>
  <div>
    <FormInput
      v-model="licenseKey"
      name="License Key"
      placeholder="Enter license key"
      :required="true"
    />
    <button @click="verifyLicense">Verify License</button>
    
    <div v-if="license">
      <h3>License Details</h3>
      <p>Status: {{ license.status }}</p>
      <p>Expires: {{ formatDate(license.expires_at) }}</p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { apiPost } from '@/utils/api'
import { formatDate } from '@/utils'

const licenseKey = ref('')
const license = ref(null)

const verifyLicense = async () => {
  try {
    const response = await apiPost('/api/license/verify', {}, {
      license_key: licenseKey.value
    })
    
    if (response.valid) {
      license.value = response.license
      window.dispatchEvent(new CustomEvent('connextSuccess', {
        detail: 'License verified successfully!'
      }))
    } else {
      window.dispatchEvent(new CustomEvent('connextError', {
        detail: 'Invalid license key'
      }))
    }
  } catch (error) {
    window.dispatchEvent(new CustomEvent('connextError', {
      detail: 'Verification failed'
    }))
  }
}
</script>
```

### File System Operations Example
```go
package main

import (
    "fmt"
    "github.com/werbot/lime/pkg/fsutil"
)

func main() {
    // Create directories
    err := fsutil.MkDirs(0755, "/path/to/dir1", "/path/to/dir2")
    if err != nil {
        fmt.Printf("Error creating directories: %v\n", err)
        return
    }
    
    // Write file
    file, err := fsutil.OpenFile("/path/to/dir1/test.txt", fsutil.FsCWTFlags, 0644)
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    
    n, err := fsutil.WriteOSFile(file, "Hello, World!")
    if err != nil {
        fmt.Printf("Error writing file: %v\n", err)
        return
    }
    
    fmt.Printf("Wrote %d bytes\n", n)
    
    // Check if file exists
    if fsutil.IsFile("/path/to/dir1/test.txt") {
        fmt.Println("File exists")
    }
}
```

### JWT Token Example
```go
package main

import (
    "fmt"
    "github.com/werbot/lime/pkg/jwtutil"
)

func main() {
    // Load keys
    err := jwtutil.LoadKeys("public.pem", "private.pem")
    if err != nil {
        fmt.Printf("Error loading keys: %v\n", err)
        return
    }
    
    // Generate token
    token, err := jwtutil.NewToken("user123", "24h", []string{"admin", "read"})
    if err != nil {
        fmt.Printf("Error generating token: %v\n", err)
        return
    }
    
    fmt.Printf("Generated token: %s\n", token)
    
    // Extract metadata
    metadata, err := jwtutil.ExtractMetadata(token)
    if err != nil {
        fmt.Printf("Error extracting metadata: %v\n", err)
        return
    }
    
    fmt.Printf("User ID: %s\n", metadata.ID)
    fmt.Printf("Expires at: %f\n", metadata.ExpiresAt)
}
```

This documentation provides a comprehensive overview of all public APIs, functions, and components in the Lime license server project. Each section includes detailed explanations, parameter descriptions, return values, and practical examples to help developers understand and use the system effectively.