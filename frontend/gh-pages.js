var ghpages = require('gh-pages');

ghpages.publish(
    'frontend/public', // path to public directory
    {
        branch: 'gh-pages',
        repo: 'https://github.com/moshval/Findlr.git', // Update to point to your repository  
        user: {
            name: 'Sheva', // update to use your name
            email: 'shevamohammad10@gmail.com' // Update to use your email
        }
    },
    () => {
        console.log('Deploy Complete!')
    }
)