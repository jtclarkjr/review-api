export interface JwtPayload {
  sub: string; // The user ID or subject of the token
  email: string; // The email of the user
  iat?: number; // Issued at timestamp
  exp?: number; // Expiration timestamp
}
