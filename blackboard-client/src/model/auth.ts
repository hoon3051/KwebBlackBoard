export interface SignInParams{
    username: string;
    password: string;
}

export interface SignUpParams{
    username: string;
    password: string;
    displayname: string;
    studentnumber: number;
    isprofessor: boolean;
}