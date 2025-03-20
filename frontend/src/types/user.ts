export interface User {
  id: string;
  name: string;
  username: string;
  email: string;
  oauth_provider: string;
  oauth_id?: string;
  avatar_url?: string;
  created_at: string;
  updated_at: string;
}
