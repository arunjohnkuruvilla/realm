import { useQuery } from "@apollo/client";
import { useCallback, useMemo, useState } from "react";
import { QuestProps, Task } from "../../../utils/consts";
import { GET_QUEST_QUERY } from "../../../utils/queries";

export const useQuests = () => {
    const [search, setSearch] = useState("");
    const { loading, data, error } = useQuery(GET_QUEST_QUERY);

    const getInitialQuestsTableData = useCallback((data: any) => {
        const formattedData = data?.map((quest: QuestProps) => {
            const taskDetails = quest.tasks.reduce((map: any, task: Task) => {
                const modMap = { ...map };

                if (task.execFinishedAt) {
                    modMap.finished += 1;
                }
                else if (task.execStartedAt) {
                    modMap.inprogress += 1;
                }
                else {
                    modMap.queued += 1;
                }

                if (new Date(task.lastModifiedAt) > new Date(modMap.lastUpdated)) {
                    modMap.lastUpdated = task.lastModifiedAt;
                }

                if (task.output !== "") {
                    modMap.outputCount += 1;
                }

                return modMap
            },
                {
                    finished: 0,
                    inprogress: 0,
                    queued: 0,
                    outputCount: 0,
                    lastUpdated: null
                }
            );

            return {
                id: quest.id,
                name: quest.name,
                tome: quest.tome.name,
                creator: quest.creator,
                ...taskDetails
            }
        });
        return formattedData.sort(function (a: any, b: any) { return new Date(b.lastUpdated).getTime() - new Date(a.lastUpdated).getTime() });
    },[]);

    const questTableData = getInitialQuestsTableData(data?.quests || []);

    const filteredData = useMemo(()=> questTableData.filter((quest: any)=> {
        const searchTerm = search.toLowerCase();
        const name = quest.name.toLowerCase();
        const tome = quest.tome.toLowerCase();

        return name.includes(searchTerm) || tome.includes(searchTerm);
    }),[search, questTableData]);

    const hasData = questTableData.length > 0;

    return {
        hasData,
        data: filteredData,
        loading,
        error,
        setSearch
    }
}
