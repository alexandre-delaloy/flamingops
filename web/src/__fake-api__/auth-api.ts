import type { User } from '../types/user';
import axios from 'axios';

import { createResourceId } from '../utils/create-resource-id';
import { JWT_EXPIRES_IN, JWT_SECRET, decode, sign } from '../utils/jwt';
import { wait } from '../utils/wait';

const users = [
    {
        id: '5e86809283e28b96d2d38537',
        avatar: '/static/mock-images/avatars/avatar-anika_visser.png',
        email: 'johndoe',
        name: 'francis  valla',
        password: 'secret',
        plan: 'Premium',
    },
];

class AuthApi {
    async login({ email, password }): Promise<any> {
        const formData = new FormData();
        formData.append('username', email);
        formData.append('password', password);
        await axios
            .post(process.env.NEXT_PUBLIC_API_HOST + '/token', formData)
            .then((res) => {
                if (res.status === 200) {
                    const accessToken = res.data.access_token;
                    window.localStorage.setItem('accessToken', accessToken);
                }
                if (res.status === 401) {
                    new Error('Please check your email and password');
                }
            });
    }

    async register({ email, name, password }): Promise<string> {
        await wait(1000);

        return new Promise((resolve, reject) => {
            try {
                // Check if a user already exists
                let user = users.find((_user) => _user.email === email);

                if (user) {
                    reject(new Error('User already exists'));
                    return;
                }

                user = {
                    id: createResourceId(),
                    avatar: null,
                    email,
                    name,
                    password,
                    plan: 'Standard',
                };

                users.push(user);

                const accessToken = sign({ userId: user.id }, JWT_SECRET, {
                    expiresIn: JWT_EXPIRES_IN,
                });
                resolve(accessToken);
            } catch (err) {
                console.error('[Auth Api]: ', err);
                reject(new Error('Internal server error'));
            }
        });
    }

    connect(): Promise<User> {
        return axios
            .get(process.env.NEXT_PUBLIC_API_HOST + '/myuser')
            .then((el) => {
                const user = el.data;
                return new Promise((resolve, reject) => {
                    if (!user) {
                        return;
                    }
                    try {
                        resolve({
                            id: user.id,
                            email: user.email,
                            username: user.username,
                            full_name: user.full_name,
                            disabled: user.disabled,
                            role_action: user.actions_labels,
                            role_label: user.role_label,
                        });
                    } catch (err) {
                        console.error('[Auth Api]: ', err);
                        reject(new Error('Internal server error'));
                    }
                });
            });
    }
}

export const authApi = new AuthApi();
