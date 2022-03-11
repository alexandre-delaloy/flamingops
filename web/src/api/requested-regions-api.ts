import axios from 'axios';

class RequestedRegionsApi {

    async getActiveRegions(
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
                .get(process.env.NEXT_PUBLIC_API_HOST + '/requested-regions', { data: body })
                .then((res) => {
                    if (res.status === 200) {
                        resolve(res.data);
                    }
                });
        });
    }

    async putActiveRegions(
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
                .put(process.env.NEXT_PUBLIC_API_HOST + '/requested-regions', body)
                .then((res) => {
                    if (res.status === 200) {
                        resolve(res.data);
                    }
                });
        });
    }

    async deleteActiveRegions(
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
                .delete(process.env.NEXT_PUBLIC_API_HOST + '/requested-regions', { data: body })
                .then((res) => {
                    if (res.status === 200) {
                        resolve(res.data);
                    }
                });
        });
    }
}
export const requestedRegionsApi = new RequestedRegionsApi();
