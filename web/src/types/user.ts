export interface User {
  id: number;
  username: string;
  full_name: string;
  email: string;
  disabled: boolean;
  role_label: string;
  role_action: string[];
}

export interface Users {
  users: User;
}
