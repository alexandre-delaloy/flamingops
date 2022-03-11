import axios from 'axios';

class SecretApi {

    async createSecret(
        ClientRequestToken: string,
        Description: string,
        Name: string,
        SecretString: string,
    ): Promise<any> {
        const body = {
            ClientRequestToken : ClientRequestToken,
            Description : Description,
            Name : Name,
            SecretString : SecretString,            
        };
        return new Promise((resolve, reject) => {
            axios
                .post(process.env.NEXT_PUBLIC_API_HOST + '/secrets', body)
                .then((res) => {
                    if (res.status === 200) {
                        resolve(res.data);
                    }
                });
        });
    }

    async modifySecrets(
        ClientRequestToken: string,
        SecretId: string,
        SecretString: string,
    ): Promise<any> {
        const body = {
            ClientRequestToken : ClientRequestToken,
            SecretId : SecretId,
            SecretString : SecretString,            
        };
        return new Promise((resolve, reject) => {
            axios
                .put(process.env.NEXT_PUBLIC_API_HOST + '/secrets/:secretName', body)
                .then((res) => {
                    if (res.status === 200) {
                        resolve(res.data);
                    }
                });
        });
    }

    async deleteSecrets(
        RecoveryWindowInDays: string,
        SecretId: string,
    ): Promise<any> {
        const body = {
            RecoveryWindowInDays : RecoveryWindowInDays,
            SecretId : SecretId,
        };
        return new Promise((resolve, reject) => {
            axios
                .delete(process.env.NEXT_PUBLIC_API_HOST + '/secrets/:secretName', { data: body })
                .then((res) => {
                    if (res.status === 200) {
                        resolve(res.data);
                    }
                });
        });
    }
}
export const secretApi = new SecretApi();
