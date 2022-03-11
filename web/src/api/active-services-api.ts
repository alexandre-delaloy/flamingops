import axios from 'axios';

class ActiveServiceApi {

    async getActiveServices(
        UserId: string,
        AwsServices: string,
        SwServices: string,
    ): Promise<any> {
        const body = {
            UserId : UserId,
            AwsServices : AwsServices,
            SwServices : SwServices,
        };
        return new Promise((resolve, reject) => {
            axios
                .get(process.env.NEXT_PUBLIC_API_HOST + '/active-services', { data: body })
                .then((res) => {
                    if (res.status === 200) {
                        resolve(res.data);
                    }
                });
        });
    }

    async putActiveServices(
        UserId: string,
        AwsServices: string,
        SwServices: string,
    ): Promise<any> {
        const body = {
            UserId : UserId,
            AwsServices : AwsServices,
            SwServices : SwServices,
        };
        return new Promise((resolve, reject) => {
            axios
                .put(process.env.NEXT_PUBLIC_API_HOST + '/active-services', body)
                .then((res) => {
                    if (res.status === 200) {
                        resolve(res.data);
                    }
                });
        });
    }

    async deleteActiveServices(
        UserId: string,
        AwsServices: string,
        SwServices: string,
    ): Promise<any> {
        const body = {
            UserId : UserId,
            AwsServices : AwsServices,
            SwServices : SwServices,
        };
        return new Promise((resolve, reject) => {
            axios
                .delete(process.env.NEXT_PUBLIC_API_HOST + '/active-services', { data: body })
                .then((res) => {
                    if (res.status === 200) {
                        resolve(res.data);
                    }
                });
        });
    }
}
export const activeServiceApi = new ActiveServiceApi();
