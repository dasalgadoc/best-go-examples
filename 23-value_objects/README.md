# 💁🏻 Explanation

In this example there are two representations:

- User without value object
- User with value object

## 📝 User without value objects

- User is a struct
- User has attributes as primitive types
- Each attribute has its own validation
- User can change its email, a new validation is needed

𝓍 __Cons:__
1. A lot of responsibilities for one struct
2. This struct is not reusable
3. If we need to change the validation, we need to change the struct (bad for SRP)

## 📝 User with value objects

- User is a struct
- User has attributes as structs
- Each struct implements its own validation for itself

✔ __Pros:__
1. Each struct has its own responsibility
2. Each struct is reusable
3. If we need to change the validation, we need to change ONE struct (good for SRP)

### 📋 Considerations

Constructor works with primitives, so if the modeling changes, we need to change on clients.
```go
func NewUserVo(id, email string, birthDate time.Time) (*UserVo, error) 
```

Mutations are safe because they work with value objects.

A email changes by another value object
```go
func (u *UserVo) UpdateEmail(newEmail string) error {
    userEmail, err := NewUserEmail(newEmail)
    if err != nil {
        return err
    }
    u.email = *userEmail
    return nil
}
```

Tell don't ask, and demeter law can be respected, but we expose getters too.