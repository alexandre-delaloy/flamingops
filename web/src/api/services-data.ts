import axios from 'axios';

class ServicesDataApi {

    async getAwsServicesData(
        UserId: string,
    ): Promise<any> {
        const body = {
            UserId : UserId,
        };
        return new Promise((resolve, reject) => {
            axios
                .get(process.env.NEXT_PUBLIC_API_HOST + '/aws-services-data', { data: body })
                .then((res) => {
                    if (res.status === 200) {
                        resolve(res.data);
                    }
                });
        });
    }

    async deleteAwsServicesData(
        UserId: string,
    ): Promise<any> {
        const body = {
            UserId : UserId,
        };
        return new Promise((resolve, reject) => {
            axios
                .delete(process.env.NEXT_PUBLIC_API_HOST + '/aws-services-data', { data: body })
                .then((res) => {
                    if (res.status === 200) {
                        resolve(res.data);
                    }
                });
        });
    }

    async getSwServicesData(
        UserId: string,
    ): Promise<any> {
        const body = {
            UserId : UserId,
        };
        return new Promise((resolve, reject) => {
            axios
                .get(process.env.NEXT_PUBLIC_API_HOST + '/sw-services-data', { data: body })
                .then((res) => {
                    if (res.status === 200) {
                        resolve(res.data);
                    }
                });
        });
    }

    async deleteSwServicesData(
        UserId: string,
    ): Promise<any> {
        const body = {
            UserId : UserId,
        };
        return new Promise((resolve, reject) => {
            axios
                .delete(process.env.NEXT_PUBLIC_API_HOST + '/sw-services-data', { data: body })
                .then((res) => {
                    if (res.status === 200) {
                        resolve(res.data);
                    }
                });
        });
    }
}
export const servicesDataApi = new ServicesDataApi();
