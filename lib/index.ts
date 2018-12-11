export interface HelloJsiiProps {
    /**
     * The message to emit when saying goodbye.
     * @default "Goodbye"
     */
    goodbyeMessage?: string;
}

/**
 * This guy knows how to say hello, in multiple langauges.
 */
export class HelloJsii {
    
    private readonly goodbyeMessage: string;
    
    constructor(props: HelloJsiiProps = {}) {
        this.goodbyeMessage = props.goodbyeMessage || 'Goodbye';
    }
    
    /**
     * Say hello to someone.
     * @param name That special someone
     */
    public sayHello(name: string) {
        return `Hello, ${name}!!`;
    }
    
    /**
     * Says goodbye, but not farewell.
     * @param times Number of times to say goodbye.
     */
    public sayGoodbye(times: number = 1) {
        let ret = '';
        for (let i = 0; i < times; ++i) {
            ret += `${this.goodbyeMessage}`;
            if (i < times - 1) {
                ret += '\n';
            }
        }
        return ret;
    }
}
