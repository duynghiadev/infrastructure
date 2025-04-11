import request from 'supertest';
import { app } from '../../app';

it('fails when an incorrect password is supplied', async () => {
    await request(app)
        .post('/api/users/signup')
        .send({
            email: "test@test.com",
            password: "password"
        })
        .expect(201)

    const response = await request(app)
        .post('/api/users/signout')
        .send({})

})