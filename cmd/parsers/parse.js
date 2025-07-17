const fs = require('fs/promises');

(async () => {
    const url = 'https://oilguide.ravenol.de/oilguideapi/en/'; // replace with actual base URL

    // Fetch all makers for rid = 1 (cars)
    const makerR = await fetch(url + 'getMakers/1');
    const maker = await makerR.json();

    const makers = await Promise.all(
        maker.response.map(async ({ label, value, image }) => {
            console.log(`starting maker ${label}`)
            const modelRes = await fetch(url + 'getModels/' + value);
            const modelJson = await modelRes.json();

            const models = await Promise.all(
                modelJson.response.map(async ({ label, value }) => {
                    console.log(`getting model ${label}`)
                    const typeRes = await fetch(url + 'getTypes/' + value);
                    const typeJson = await typeRes.json();
                    console.log(`got types ${JSON.stringify(typeJson)}`)

                    return {
                        label,
                        rid: value,
                        types: typeJson.response,
                    };
                })
            );

            return {
                label,
                rid: value,
                image,
                models,
            };
        })
    );

    const cars = {
        label: 'cars',
        rid: 1,
        makers,
    };
    // Write to cars.json
    try {
        await fs.writeFile('cars.json', JSON.stringify(cars, null, 2));
        console.log('✅ cars.json saved successfully.');
    } catch (err) {
        console.error('❌ Failed to save cars.json:', err);
    }
})();