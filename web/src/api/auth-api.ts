// import type { User } from '../types/user';
// import axios from 'axios';

// import { wait } from '../utils/wait';


// class AuthApi {
//     async login({ username, mail, password }): Promise<any> {
//         const formData = new FormData();
//         formData.append('username', username);
//         formData.append('mail', mail);
//         formData.append('password', password);
//         await axios
//             .post(process.env.NEXT_PUBLIC_API_HOST + '/login', formData)
//             .then((res) => {
//                 if (res.status === 200) {
//                     const accessToken = res.data.access_token;
//                     window.localStorage.setItem('accessToken', accessToken);
//                 }
//                 if (res.status === 401) {
//                     new Error('Please check your email and password');
//                 }
//             });
//     }

//     async register({ mail, username, password }): Promise<any> {
//         await wait(1000);

//         const formData = new FormData();
//         formData.append('username', username);
//         formData.append('mail', mail);
//         formData.append('password', password);
//         await axios
//             .post(process.env.NEXT_PUBLIC_API_HOST + '/users', formData)
//             .then((res) => {
//                 if (res.status === 200) {
//                     return res.data;
//                 }
//                 if (res.status === 401) {
//                     new Error('Please check your email and password');
//                 }
//             });
//     }

//     connect(): Promise<User> {
//         return axios
//             .get(process.env.NEXT_PUBLIC_API_HOST + '/myuser')
//             .then((el) => {
//                 const user = el.data;
//                 return new Promise((resolve, reject) => {
//                     if (!user) {
//                         return;
//                     }
//                     try {
//                         resolve({
//                             id: user.id,
//                             email: user.email,
//                             username: user.username,
//                         });
//                     } catch (err) {
//                         console.error('[Auth Api]: ', err);
//                         reject(new Error('Internal server error'));
//                     }
//                 });
//             });
//     }
// }

// export const authApi = new AuthApi();
import type { User } from '../types/user';
import { createResourceId } from '../utils/create-resource-id';
import { sign, decode, JWT_SECRET, JWT_EXPIRES_IN } from '../utils/jwt';
import { wait } from '../utils/wait';

const users = [
  {
    id: '5e86809283e28b96d2d38537',
    avatar: '/static/mock-images/avatars/avatar-anika_visser.png',
    email: 'demo@devias.io',
    name: 'Anika Visser',
    password: 'Password123!',
    plan: 'Premium'
  }
];

class AuthApi {
  async login({ email, password }): Promise<string> {
    await wait(500);

    return new Promise((resolve, reject) => {
      try {
        // Find the user
        const user = users.find((_user) => _user.email === email);

        if (!user || (user.password !== password)) {
          reject(new Error('Please check your email and password'));
          return;
        }

        // Create the access token
        const accessToken = sign(
          { userId: user.id },
          JWT_SECRET,
          { expiresIn: JWT_EXPIRES_IN }
        );

        resolve(accessToken);
      } catch (err) {
        console.error('[Auth Api]: ', err);
        reject(new Error('Internal server error'));
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
          id: 'oui',
          avatar: null,
          email,
          name,
          password,
          plan: 'Standard'
        };

        users.push(user);

        const accessToken = sign(
          { userId: user.id },
          JWT_SECRET,
          { expiresIn: JWT_EXPIRES_IN }
        );

        resolve(accessToken);
      } catch (err) {
        console.error('[Auth Api]: ', err);
        reject(new Error('Internal server error'));
      }
    });
  }

  me(accessToken): Promise<User> {
    return new Promise((resolve, reject) => {
      try {
        // Decode access token
        const { userId } = decode(accessToken) as any;

        // Find the user
        const user = users.find((_user) => _user.id === userId);

        if (!user) {
          reject(new Error('Invalid authorization token'));
          return;
        }

        resolve({
            id: 43,
            email: user.email,
            username: user.name,
        });
      } catch (err) {
        console.error('[Auth Api]: ', err);
        reject(new Error('Internal server error'));
      }
    });
  }
}
export const authApi = new AuthApi();