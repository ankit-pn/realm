import { useFormik } from "formik";
import { faker } from '@faker-js/faker';
import React, { useState } from "react";

import { FormSteps } from "../../../components/form-steps";
import TomeStepWrapper from "./TomeStepWrapper";
import FinalizeStep from "./FinalizeStep";
import BeaconStepWrapper from "./BeaconStepWrapper";
import { useSubmitQuest } from "../hooks/useSubmitQuest";

const QuestForm = () => {
    const [currStep, setCurrStep] = useState<number>(0);
    const { submitQuest } = useSubmitQuest();
    const placeholderTitle = faker.music.songName();


    const steps = [
        { name: 'Select agent beacons', description: 'Step 1', href: '#', step: 0 },
        { name: 'Select a tome', description: 'Step 2', href: '#', step: 1 },
        { name: 'Confirm quest details', description: 'Step 3', href: '#', step: 2 },
    ];

    const formik = useFormik({
        initialValues: {
            name: placeholderTitle,
            tome: null,
            params: [],
            beacons: [],
        },
        onSubmit: (values: any) => submitQuest(values),
    });

    function getStepView(step: number) {
        switch (step) {
            case 0:
                return <BeaconStepWrapper setCurrStep={setCurrStep} formik={formik} />
            case 1:
                return <TomeStepWrapper setCurrStep={setCurrStep} formik={formik} />
            case 2:
                return <FinalizeStep setCurrStep={setCurrStep} formik={formik} />
            default:
                return <div>An error has occured</div>;
        }
    }

    return (
        <form
            id='create-quest-form'
            className="py-6"
        >
            <div className="grid grid-cols-12">
                <div className=" col-span-3">
                    <FormSteps currStep={currStep} steps={steps} />
                </div>
                <div className="col-span-9">
                    {getStepView(currStep)}
                </div>
            </div>
        </form>
    );
}
export default QuestForm;
