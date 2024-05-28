export const getSpecificSchool = defineCachedFunction(
  async (npsn: npsn) => {
    const data = await useStorage("assets:server").getItem("schools.json");

    const parsedData = schoolsDataScheme.parse(data);

    return parsedData.find((school) => school.npsn === npsn) ?? null;
  },
  {
    maxAge: 60 * 60 * 2,
    name: "npsnKey",
    getKey: (npsn: number) => npsn,
  },
);
