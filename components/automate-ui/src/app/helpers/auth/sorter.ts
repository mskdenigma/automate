
export class ChefSorters {

    // naturalSort takes in an Array of Objects and returns an Array of Objects
    // sorted 'naturally' by the property chosen.

    public static naturalSort( input: Array<any>, property: Array<string> | string ): Array<any> {
        // per @msorens https://github.com/chef/a2/pull/4434
        // Stable sort so 'a' always comes before 'A'.
        console.log('********');
        console.log( property );
        const opts = { numeric: true, sensitivity: 'base' };

        // propComparator function idea from Dave Newton
        // https://stackoverflow.com/questions
        const propComparator = (propName) => {
            return (a, b) => a[propName].localeCompare(b[propName], undefined, { opts }) ||
                      a[propName].localeCompare(b[propName], undefined, { numeric: true });
        };

        const dualPropComparator = (propOne, propTwo) => {
            return (a, b) => a[propOne].localeCompare(b[propOne], undefined, { opts }) ||
                      a[propOne].localeCompare(b[propOne], undefined, { numeric: true }) ||
                      a[propTwo].localeCompare(b[propTwo], undefined, opts);
        };

        if ( property instanceof Array ) {
            console.log('array');
            return input.sort( dualPropComparator(property[0], property[1]) );
        } else {
            console.log('string');
            return input.sort( propComparator(property) );
        }

    }
}